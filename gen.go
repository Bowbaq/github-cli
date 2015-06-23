package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

type service struct {
	Name        string
	SubCommands []command
}

type command struct {
	Method method
	Tmpl   *template.Template
	flags  []flag
}

func (c command) Body() string {
	var buf bytes.Buffer

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed template instanciation", r)
			buf = bytes.Buffer{}
		}
	}()

	err := c.Tmpl.Execute(&buf, c)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (c command) Flags() []flag {
	for _, arg := range c.Method.Args {
		switch {
		// Int and string, *os.File don't generate flags
		case arg.Typ == "int":
		case arg.Typ == "string":
		case arg.Typ == "*os.File":
		case arg.Typ == "bool":
			fallthrough
		case arg.Typ == "[]string":
			fallthrough
		case arg.Typ == "time.Time":
			c.flags = append(c.flags, flag{Typ: arg.Typ, Name: arg.Name})
		case strings.HasPrefix(arg.Typ, "*github."):
			typeName := strings.TrimPrefix(arg.Typ, "*github.")
			c.flags = append(c.flags, flagSet(typeName, types[typeName])...)
		default:
			log.Println("unimplemented arg type: ", arg.Typ)
		}
	}

	return c.flags
}

func (c command) Usage() string {
	var usage bytes.Buffer
	usage.WriteString(dasherize(c.Method.Name) + " ")
	for _, arg := range c.Method.Args {
		if arg.Typ != "string" && arg.Typ != "int" && arg.Typ != "*os.File" {
			continue
		}
		usage.WriteString("<" + dasherize(arg.Name) + "> ")
	}

	return strings.TrimSpace(usage.String())
}

func (c command) UsageCount() int {
	var count int

	for _, arg := range c.Method.Args {
		if arg.Typ == "string" || arg.Typ == "int" || arg.Typ == "*os.File" {
			count += 1
		}
	}

	return count
}

func (c command) SetupArgs() string {
	var setup []string
	for i, arg := range c.Method.Args {
		switch {
		case arg.Typ == "int":
			setup = append(setup, fmt.Sprintf("%s, err := strconv.Atoi(c.Args().Get(%d))", arg.Name, i), "check(err)")
		case arg.Typ == "bool":
			setup = append(setup, fmt.Sprintf(`%s := c.Bool("%s")`+"\n", arg.Name, dasherize(arg.Name)))
		case arg.Typ == "string":
			setup = append(setup, fmt.Sprintf("%s := c.Args().Get(%d)", arg.Name, i))
		case arg.Typ == "*os.File":
			setup = append(setup, fmt.Sprintf("%s, err := os.Open(c.Args().Get(%d))", arg.Name, i), "check(err)")
		case strings.HasPrefix(arg.Typ, "*github."):
			typeName := strings.TrimPrefix(arg.Typ, "*github.")
			typeInfo := types[typeName]
			setup = append(setup, fmt.Sprintf("%s := &github.%s{", arg.Name, typeName))
			for _, flag := range flagSet(typeName, typeInfo) {
				if _, ok := typeInfo[flag.Name]; !ok {
					continue // Ignore flags that didn't come from the type definition
				}
				setup = append(setup, fmt.Sprintf("%s: %s,", flag.Name, flag.Accessor()))
			}
			setup = append(setup, "}")
		default:
			isFlag := false
			for _, flag := range c.Flags() {
				if flag.Name == arg.Name {
					setup = append(setup, fmt.Sprintf("%s := %s", arg.Name, flag.Accessor()))
					isFlag = true
				}
			}
			if !isFlag {
				if isExported(arg.Typ) {
					setup = append(setup, fmt.Sprintf("var %s %s", arg.Name, arg.Typ))
				} else {
					panic(fmt.Sprintf("%s - missing argument %s %s", c.Method, arg.Name, arg.Typ))
				}
			}
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

var (
	methods []method
	types   map[string]map[string]flag
)

func main() {
	methods, types = analyseAST()

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

	// f, err := os.Create(path.Join("cmd", "github", "README.md"))
	// check(err)

	// fmt.Fprintln(f, "# github-cli")

	// fmt.Fprintln(f, "## Implemented")
	// fmt.Fprintln(f, "```go")
	// done := 0
	// for _, pair := range sortMapByValue(implemented) {
	// 	fmt.Fprintln(f, pair.Key)
	// 	for _, sig := range implemented[pair.Key] {
	// 		done += 1
	// 		fmt.Fprintln(f, "  ", sig)
	// 	}
	// }
	// fmt.Fprintln(f, "```")

	// fmt.Fprintln(f, "## Unimplemented")
	// fmt.Fprintln(f, "```go")
	// for _, pair := range sortMapByValue(unimplemented) {
	// 	fmt.Fprintln(f, pair.Key)
	// 	for _, sig := range unimplemented[pair.Key] {
	// 		fmt.Fprintln(f, "  ", sig)
	// 	}
	// }
	// fmt.Fprintln(f, "```")

	// fmt.Println("Implemented", done, "/", len(methods))
}

func toSubCommand(m method) *command {
	cmd := &command{Method: m, Tmpl: notImplementedTmpl}
	if isSimpleListMethod(m) {
		cmd.Tmpl = listTmpl
	} else {
		if len(m.Returns) <= 3 {
			cmd.Tmpl = singleTmpl
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

func flagSet(typeName string, typeInfo map[string]flag) []flag {
	var flags []flag
	for _, f := range typeInfo {
		if strings.HasSuffix(f.Name, "URL") {
			continue
		}
		if f.Name == "TextMatches" {
			continue
		}

		switch {
		case f.Typ == "int":
			fallthrough
		case f.Typ == "*int":
			fallthrough
		case f.Typ == "bool":
			fallthrough
		case f.Typ == "*bool":
			fallthrough
		case f.Typ == "string":
			fallthrough
		case f.Typ == "*string":
			fallthrough
		case f.Typ == "[]string":
			fallthrough
		case f.Typ == "*[]string":
			fallthrough
		case f.Typ == "time.Time":
			fallthrough
		case f.Typ == "*time.Time":
			fallthrough
		case f.Typ == "*github.Timestamp":
			flags = append(flags, f)
		default:
			if strings.HasPrefix(f.Typ, "github.") || strings.HasPrefix(f.Typ, "*github.") {
				subTypeName := strings.TrimPrefix(strings.TrimPrefix(f.Typ, "*"), "github.")
				if subTypeName == typeName {
					continue
				}

				subTypeInfo := types[subTypeName]
				subFlags := flagSet(subTypeName, subTypeInfo)
				if f.Name == "" {
					flags = append(flags, subFlags...)
				} else {
					for _, sf := range subFlags {
						flags = append(flags, flag{Typ: sf.Typ, Name: dasherize(f.Name + "-" + sf.Name), Usage: sf.Usage})
					}
				}
			} else {
				log.Println("unimplemented flag type", f.Typ)
			}
		}
	}
	if typeName == "ListOptions" {
		flags = append(flags, flag{Typ: "bool", Name: "all", Usage: `For paginated result sets, fetch all remaining pages starting at "page"`})
	}

	return flags
}

func isSimpleListMethod(m method) bool {
	if !strings.HasPrefix(m.Returns[0], "[]") {
		return false
	}

	re := regexp.MustCompile(".*List.*Options")
	for _, arg := range m.Args {
		typeName := strings.TrimPrefix(arg.Typ, "*github.")
		if re.MatchString(typeName) {
			if typeName == "ListOptions" {
				return true
			}
			for _, f := range types[typeName] {
				if f.Typ == "github.ListOptions" {
					return true
				}
			}
		}
	}
	return false
}

func isExported(typ string) bool {
	parts := strings.Split(typ, ".")
	last := parts[len(parts)-1]
	return unicode.IsUpper(rune(last[0]))
}
