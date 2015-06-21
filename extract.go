package main

import (
	"fmt"
	"go/ast"
	"strings"

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

func (m method) signature() string {
	var strargs []string
	for _, arg := range m.args {
		strargs = append(strargs, arg.typ)
	}

	return "(" + strings.Join(strargs, ", ") + ")"
}

func extractServiceMethods() (methods []method) {
	var conf loader.Config

	conf.Import("github.com/google/go-github/github")
	prog, err := conf.Load()
	if err != nil {
		return methods
	}

	pkg := prog.Package("github.com/google/go-github/github")
	for _, f := range pkg.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			if method := toServiceMethod(pkg, node); method != nil {
				methods = append(methods, *method)
			}

			return true
		})
	}

	return methods
}

func toServiceMethod(pkg *loader.PackageInfo, n ast.Node) *method {
	// Find function declarations
	decl, ok := n.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	// Discard functions, we only want methods
	if decl.Recv == nil {
		return nil
	}

	// Only want methods on a *Service type
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
	// Extract (name, type) pairs of method arguments
	for _, arg := range decl.Type.Params.List {
		a := argument{
			name: arg.Names[0].Name,
			typ:  strings.Replace(pkg.Info.TypeOf(arg.Type).String(), "github.com/google/go-github/", "", -1),
		}
		m.args = append(m.args, a)
	}
	// Extract method return types
	for _, ret := range decl.Type.Results.List {
		m.returns = append(
			m.returns,
			strings.Replace(pkg.Info.TypeOf(ret.Type).String(), "github.com/google/go-github/", "", -1),
		)
	}

	return m
}
