package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"strings"

	"golang.org/x/tools/go/loader"
)

type argument struct {
	Name string
	Typ  string
}

func (a argument) String() string {
	return a.Name + " " + a.Typ
}

type method struct {
	Service     string
	Name        string
	Description string
	Args        []argument
	Returns     []string
}

func (m method) String() string {
	var strargs []string
	for _, arg := range m.Args {
		strargs = append(strargs, arg.String())
	}

	return fmt.Sprintf("%s.%s(%s) (%s)", m.Service, m.Name, strings.Join(strargs, ", "), strings.Join(m.Returns, ", "))
}

func (m method) Usage() string {
	return doc.Synopsis(m.Description)
}

func (m method) signature() string {
	var strargs []string
	for _, arg := range m.Args {
		strargs = append(strargs, arg.Typ)
	}

	return "(" + strings.Join(strargs, ", ") + ")"
}

func extractServiceMethods() (methods []method) {
	var conf loader.Config

	conf.ParserMode |= parser.ParseComments
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

	// Only want exported methods
	if !ast.IsExported(decl.Name.String()) {
		return nil
	}

	m := &method{
		Service:     ident.Name,
		Name:        decl.Name.String(),
		Description: formatDescription(decl.Name.String(), decl.Doc.Text()),
	}

	// Extract (name, type) pairs of method arguments
	for _, arg := range decl.Type.Params.List {
		typ := strings.Replace(pkg.Info.TypeOf(arg.Type).String(), "github.com/google/go-github/", "", -1)
		for _, name := range arg.Names {
			m.Args = append(m.Args, argument{
				Name: name.Name,
				Typ:  typ,
			})
		}
	}
	// Extract method return types
	for _, ret := range decl.Type.Results.List {
		m.Returns = append(
			m.Returns,
			strings.Replace(pkg.Info.TypeOf(ret.Type).String(), "github.com/google/go-github/", "", -1),
		)
	}

	return m
}

// Fix indentation in CLI output
func formatDescription(methodName, desc string) string {
	desc = strings.Replace(desc, methodName, dasherize(methodName), -1)

	lines := strings.Split(desc, "\n")
	if len(lines) <= 1 {
		return desc
	}

	for i, line := range lines[1:] {
		if len(line) > 0 {
			lines[i+1] = "   " + line
		}
	}

	return strings.Join(lines, "\n")
}
