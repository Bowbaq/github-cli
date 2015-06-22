package main

import (
	"bytes"
	"fmt"
	"log"
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

func (c command) Flags() (flags []string) {
	for _, arg := range c.Method.Args {
		switch {
		case arg.Typ == "bool":
			flags = append(flags, fmt.Sprintf(`cli.BoolFlag{Name:  "%s"}`, dasherize(arg.Name)))
		case arg.Typ == "*github.ListOptions":
			flags = append(flags, `cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"}`)
			flags = append(flags, `cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"}`)
			flags = append(flags, `cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"}`)
		case arg.Typ == "int":
		case arg.Typ == "string":
		default:
			log.Println("Unimplemented flag type: ", arg.Typ)
		}
	}

	return flags
}

func (c command) Usage() string {
	var usage bytes.Buffer
	usage.WriteString(dasherize(c.Method.Name) + " ")
	for _, arg := range c.Method.Args {
		if arg.Typ != "string" && arg.Typ != "int" {
			break
		}
		usage.WriteString("<" + dasherize(arg.Name) + "> ")
	}

	return strings.TrimSpace(usage.String())
}

func (c command) SetupArgs() string {
	var setup []string
	for i, arg := range c.Method.Args {
		switch arg.Typ {
		case "string":
			setup = append(setup, fmt.Sprintf("%s := c.Args().Get(%d)", arg.Name, i))
		case "int":
			setup = append(setup, fmt.Sprintf("%s, err := strconv.Atoi(c.Args().Get(%d))", arg.Name, i), "check(err)")
		case "bool":
			setup = append(setup, fmt.Sprintf(`%s := c.Bool("%s")`+"\n", arg.Name, dasherize(arg.Name)))
		case "*github.ListOptions":
			setup = append(setup, arg.Name+` := &github.ListOptions{
        Page: c.Int("page"),
        PerPage: c.Int("page-size"),
      }`)
		}
	}

	return strings.Join(setup, "\n")
}

func (c command) ArgList() string {
	var list []string
	for _, arg := range c.Method.Args {
		list = append(list, arg.Name)
	}

	return strings.Join(list, ", ")
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
	fmt.Fprintln(f, "```go")
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
	fmt.Fprintln(f, "```go")
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
	case "(int, *github.ListOptions)":
		fallthrough
	case "(string, *github.ListOptions)":
		fallthrough
	case "(string, bool, *github.ListOptions)":
		fallthrough
	case "(string, string, *github.ListOptions)":
		fallthrough
	case "(string, string, int, *github.ListOptions)":
		fallthrough
	case "(string, string, string, *github.ListOptions)":
		if strings.HasPrefix(m.Returns[0], "[]") {
			cmd.Tmpl = listTmpl
		}

	case "()":
		fallthrough
	case "(int)":
		fallthrough
	case "(string)":
		fallthrough
	case "(string, string)":
		fallthrough
	case "(string, int)":
		fallthrough
	case "(int, string)":
		fallthrough
	case "(string, string, int)":
		fallthrough
	case "(int, string, string)":
		fallthrough
	case "(string, string, string)":
		fallthrough
	case "(string, string, int, string)":
		fallthrough
	case "(string, string, string, bool)":
		cmd.Tmpl = singleTmpl
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
