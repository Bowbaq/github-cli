package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"go/ast"

	"golang.org/x/tools/go/loader"
)

type argument struct {
	name string
	typ  string
}

func (a argument) String() string {
	return a.name + " " + a.typ
}

type method struct {
	service string
	name    string
	args    []argument
	returns []string
}

func (m method) String() string {
	var strargs []string
	for _, arg := range m.args {
		strargs = append(strargs, arg.String())
	}

	return fmt.Sprintf("%s.%s(%s)", m.service, m.name, strings.Join(strargs, ", "))
}

func main() {
	var conf loader.Config

	conf.Import("github.com/google/go-github/github")

	prog, err := conf.Load()
	check(err)
	pkg := prog.Package("github.com/google/go-github/github")

	var methods []*method
	for _, f := range pkg.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if method := toServiceMethod(pkg, n); method != nil {
				methods = append(methods, method)
			}

			return true
		})
	}

	impls := make(map[string][]string)
	for _, method := range methods {
		switch {
		case len(method.args) == 1 && method.args[0].typ == "*github.ListOptions":
			impls[method.service] = append(impls[method.service], simpleListImpl(method))
		}
	}

	for service, commands := range impls {
		f, err := os.Create(path.Join("cmd", "github", camelcase(service)+".go"))
		check(err)

		subCommandTmpl.Execute(f, map[string]interface{}{
			"Service": service,
			"Imports": []string{
				"fmt",
				"github.com/codegangsta/cli",
				"github.com/google/go-github/github",
			},
			"Commands": commands,
		})

		f.Close()
	}

	check(exec.Command("goimports", "-w", "cmd/github").Run())
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func toServiceMethod(pkg *loader.PackageInfo, n ast.Node) *method {
	decl, ok := n.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	if decl.Recv == nil {
		return nil
	}

	recv, ok := decl.Recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return nil
	}

	ident, ok := recv.X.(*ast.Ident)
	if !ok {
		return nil
	}

	if !strings.HasSuffix(ident.Name, "Service") {
		return nil
	}

	m := &method{
		service: ident.Name,
		name:    decl.Name.String(),
	}
	for _, arg := range decl.Type.Params.List {
		a := argument{
			name: arg.Names[0].Name,
			typ:  strings.Replace(pkg.Info.TypeOf(arg.Type).String(), "github.com/google/go-github/", "", -1),
		}
		m.args = append(m.args, a)
	}
	for _, ret := range decl.Type.Results.List {
		retType := strings.Replace(pkg.Info.TypeOf(ret.Type).String(), "github.com/google/go-github/", "", -1)
		m.returns = append(m.returns, retType)
	}

	return m
}
