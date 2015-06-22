package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"text/template"
)

type service struct {
	Name        string
	SubCommands []command
}

type command struct {
	Method method
	Tmpl   *template.Template
}

func (c command) Body() string {
	var buf bytes.Buffer
	check(c.Tmpl.Execute(&buf, c))
	return buf.String()
}

func main() {
	methods := extractServiceMethods()

	services := make(map[string]*service)

	implemented := make(map[string][]string)
	unimplemented := make(map[string][]string)

	for _, method := range methods {
		if _, ok := services[method.Service]; !ok {
			services[method.Service] = &service{Name: method.Service}
		}

		subCommand := toSubCommand(method)

		services[method.Service].SubCommands = append(services[method.Service].SubCommands, *subCommand)

		if subCommand.Tmpl == notImplementedTmpl {
			unimplemented[method.signature()] = append(unimplemented[method.signature()], method.String())
		} else {
			implemented[method.signature()] = append(implemented[method.signature()], method.String())
		}
	}

	for name, service := range services {
		f, err := os.Create(path.Join("cmd", "github", camelcase(name)+".go"))
		check(err)

		check(serviceTmpl.Execute(f, service))
		f.Close()
	}

	check(exec.Command("goimports", "-w", "cmd/github").Run())

	f, err := os.Create(path.Join("cmd", "github", "README.md"))
	check(err)

	fmt.Fprintln(f, "# github-cli")

	fmt.Fprintln(f, "## Implemented")
	fmt.Fprintln(f, "```")
	done := 0
	for _, pair := range sortMapByValue(implemented) {
		fmt.Fprintln(f, pair.Key)
		for _, sig := range implemented[pair.Key] {
			done += 1
			fmt.Fprintln(f, "  ", sig)
		}
	}
	fmt.Fprintln(f, "```")

	fmt.Fprintln(f, "## Unimplemented")
	fmt.Fprintln(f, "```")
	for _, pair := range sortMapByValue(unimplemented) {
		fmt.Fprintln(f, pair.Key)
		for _, sig := range unimplemented[pair.Key] {
			fmt.Fprintln(f, "  ", sig)
		}
	}
	fmt.Fprintln(f, "```")

	fmt.Println("Implemented", done, "/", len(methods))
}

func toSubCommand(m method) *command {
	cmd := &command{Method: m, Tmpl: notImplementedTmpl}
	switch m.signature() {
	case "(*github.ListOptions)":
		fallthrough
	case "(string, *github.ListOptions)":
		fallthrough
	case "(string, string, *github.ListOptions)":
		fallthrough
	case "(string, string, int, *github.ListOptions)":
		fallthrough
	case "(string, string, string, *github.ListOptions)":
		if strings.HasPrefix(m.Returns[0], "[]") {
			cmd.Tmpl = listTmpl
		}
	}

	return cmd
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string][]string) PairList {
	p := make(PairList, 0)
	for k, v := range m {
		p = append(p, Pair{k, len(v)})
	}
	sort.Sort(p)
	return p
}
