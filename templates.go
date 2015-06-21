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

var notImplementedTmpl = template.Must(template.New("not-implemented").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Name | dasherize}}",
  Action: func(c *cli.Context) {
    fmt.Println("Not implemented")
    os.Exit(1)
  },
},`))
