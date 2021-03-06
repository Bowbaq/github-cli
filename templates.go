package main

import (
	"regexp"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"dasherize": dasherize,
	"pointer": func(name string) string {
		return strings.TrimSuffix(name, "Service")
	},
	"sub": func(a, b int) int {
		return a - b
	},
}

func dasherize(name string) string {
	re := regexp.MustCompile("[a-z][A-Z]")
	return strings.ToLower(re.ReplaceAllStringFunc(name, func(s string) string {
		return s[:1] + "-" + s[1:]
	}))
}

func camelcase(name string) string {
	re := regexp.MustCompile("[a-z][A-Z]")
	return strings.ToLower(re.ReplaceAllStringFunc(name, func(s string) string {
		return s[:1] + "_" + s[1:]
	}))
}

var serviceTmpl = template.Must(template.New("service").Funcs(funcMap).Parse(`
package main

import "github.com/google/go-github/github"

var {{.Name}} = cli.Command{
  Name:        "{{.Name | pointer | dasherize}}",
  HideHelp:    true,
  Action:      fixHelp,
  Subcommands: []cli.Command{
  {{range .SubCommands}}{{.Body}}{{end}}
  },
}

func init() {
  app.cli.Commands = append(app.cli.Commands, {{.Name}})
}
`))

var singleTmpl = template.Must(template.New("single").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Method.Name | dasherize}}",
  Usage: ` + "`" + `{{.Method.Usage}}` + "`" + `,
  Description: ` + "`" + `{{.Method.Description}}` + "`" + `,
  Flags: []cli.Flag{
    {{range .Flags}}{{.Declaration}},
    {{end}}
  },
  Action: func(c *cli.Context) { {{if gt .UsageCount 0}}
    if len(c.Args()) < {{.UsageCount}} {
      showHelp(c, "{{.Method.Name | dasherize}}", "{{.Usage}}")
    }

    {{end}}
    {{.SetupArgs}}

    {{if eq (len .Method.Returns) 3}}
    result, res, err := app.gh.{{.Method.Service | pointer}}.{{.Method.Name}}({{.ArgList}})
    checkResponse(res.Response, err)
    fmt.Printf("%# v", pretty.Formatter(result))
    {{else}}
    {{if eq (index .Method.Returns 0) "*github.Response"}}
    res, err := app.gh.{{.Method.Service | pointer}}.{{.Method.Name}}({{.ArgList}})
    checkResponse(res.Response, err)
    {{else}}
    _, err := app.gh.{{.Method.Service | pointer}}.{{.Method.Name}}({{.ArgList}})
    check(err)
    {{end}}
    {{end}}
  },
},`))

var listTmpl = template.Must(template.New("list").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Method.Name | dasherize}}",
  Usage: ` + "`" + `{{.Method.Usage}}` + "`" + `,
  Description: ` + "`" + `{{.Method.Description}}` + "`" + `,
  Flags: []cli.Flag{
    {{range .Flags}}{{.Declaration}},
    {{end}}
  },
  Action: func(c *cli.Context) { {{if gt .UsageCount 0}}
    if len(c.Args()) < {{.UsageCount}} {
      showHelp(c, "{{.Method.Name | dasherize}}", "{{.Usage}}")
    }

    {{end}}
    {{.SetupArgs}}

    var items {{index .Method.Returns 0}}

    for {
      page, res, err := app.gh.{{.Method.Service | pointer}}.{{.Method.Name}}({{.ArgList}})
      checkResponse(res.Response, err)

      items = append(items, page...)
      if res.NextPage == 0 || !c.Bool("all") {
        break
      }
      opt.Page = res.NextPage
    }

    fmt.Printf("%# v", pretty.Formatter(items))
  },
},`))

var notImplementedTmpl = template.Must(template.New("not-implemented").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Method.Name | dasherize}}",
  Usage: "not implemented",
  Action: func(c *cli.Context) {
    fatalln("Not implemented")
  },
},`))
