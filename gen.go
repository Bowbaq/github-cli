package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"
	"text/template"
)

type service struct {
	Name        string
	SubCommands []command
}

type command struct {
	Service    string
	Name       string
	ReturnType string
	Tmpl       *template.Template
}

func (c command) Body() string {
	var buf bytes.Buffer
	check(c.Tmpl.Execute(&buf, c))
	return buf.String()
}

func main() {
	methods := extractServiceMethods()

	services := make(map[string]*service)
	unimplemented := make(map[string][]string)
	for _, method := range methods {
		if _, ok := services[method.service]; !ok {
			services[method.service] = &service{Name: method.service}
		}

		if subCommand := toSubCommand(method); subCommand != nil {
			services[method.service].SubCommands = append(services[method.service].SubCommands, *subCommand)
		} else {
			unimplemented[method.signature()] = append(unimplemented[method.signature()], method.String())
		}
	}

	for name, service := range services {
		f, err := os.Create(path.Join("cmd", "github", camelcase(name)+".go"))
		check(err)

		check(serviceTmpl.Execute(f, service))
		f.Close()
	}

	check(exec.Command("goimports", "-w", "cmd/github").Run())

	sorted := sortMapByValue(unimplemented)
	missing := 0
	for _, pair := range sorted {
		fmt.Println(pair.Key)
		for _, sig := range unimplemented[pair.Key] {
			missing += 1
			fmt.Println("  ", sig)
		}
	}

	fmt.Println("Implemented", len(methods)-missing, "/", len(methods))

}

func toSubCommand(m method) *command {
	var cmd *command
	switch m.signature() {
	case "(*github.ListOptions)":
		cmd = &command{Tmpl: simpleListTmpl}
	default:
		cmd = &command{Tmpl: notImplementedTmpl}
	}
	if cmd == nil {
		return nil
	}

	cmd.Service = m.service
	cmd.Name = m.name
	cmd.ReturnType = m.returns[0]

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
