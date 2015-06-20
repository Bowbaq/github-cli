package main

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"dasherize": dasherize,
	"pointer": func(name string) string {
		return strings.TrimSuffix(name, "Service")
	},
}

var subCommandTmpl = template.Must(template.New("subcommand").Funcs(funcMap).Parse(`
package main

import (
{{range .Imports}}"{{.}}"
{{end}}
)

var {{.Service}} = cli.Command{
  Name:        "{{.Service | pointer | dasherize}}",
  HideHelp:    true,
  Action:      fixHelp,
  Subcommands: []cli.Command{
  {{range .Commands}}{{.}}{{end}}
  },
}

func init() {
  app.cli.Commands = append(app.cli.Commands, {{.Service}})
}
`))

var simpleListTmpl = template.Must(template.New("simple-list").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Name | dasherize}}",
  Flags: []cli.Flag{
    cli.BoolFlag{
      Name:  "all, a",
      Usage: "fetch all the pages",
    },
    cli.IntFlag{
      Name:  "page, p",
      Value: 0,
      Usage: "fetch this specific page",
    },
    cli.IntFlag{
      Name:  "page-size, ps",
      Value: 30,
      Usage: "fetch <page-size> items per page",
    },
  },
  Action: func(c *cli.Context) {
    var items {{.ReturnType}}

    opt := &github.ListOptions{
      Page: c.Int("page"),
      PerPage: c.Int("page-size"),
    }

    for {
      page, res, err := app.gh.{{.Service | pointer}}.{{.Name}}(opt)
      checkResponse(res.Response, err)

      items = append(items, page...)
      if res.NextPage == 0 || !c.Bool("all") {
        break
      }
      opt.Page = res.NextPage
    }

    for _, item := range items {
      fmt.Println(item)
    }
  },
},`))

func simpleListImpl(m *method) string {
	var buf bytes.Buffer
	check(simpleListTmpl.Execute(&buf, map[string]string{
		"Service":    m.service,
		"Name":       m.name,
		"ReturnType": m.returns[0],
	}))

	return buf.String()
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
