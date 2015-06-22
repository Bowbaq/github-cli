package main

import (
	"fmt"
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
	"usage": func(m method) string {
		var strargs []string
		for _, arg := range m.Args {
			if strings.HasSuffix(arg.Typ, "Options") {
				break
			}
			strargs = append(strargs, "<"+arg.Name+">")
		}

		return strings.Join(strargs, " ")
	},
	"ctxargs": func(m method) string {
		var ctxargs []string
		for i, arg := range m.Args {
			if strings.HasSuffix(arg.Typ, "Options") {
				break
			}
			ctxarg := fmt.Sprintf("c.Args().Get(%d)", i)
			if arg.Typ == "int" {
				ctxarg = "parseInt(" + ctxarg + ")"
			}
			ctxargs = append(ctxargs, ctxarg)
		}

		return strings.Join(ctxargs, ", ")
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

var listTmpl = template.Must(template.New("simple-list").Funcs(funcMap).Parse(
	`cli.Command{
  Name:  "{{.Method.Name | dasherize}}",
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
  Action: func(c *cli.Context) { {{if gt (len .Method.Args) 1}}
    if len(c.Args()) < {{sub (len .Method.Args) 1}} {
      fatalln("Usage: " + c.App.Name + " {{.Method.Name | dasherize}} {{.Method | usage}}")
    }

    {{end}}
    var items {{index .Method.Returns 0}}

    opt := &github.ListOptions{
      Page: c.Int("page"),
      PerPage: c.Int("page-size"),
    }

    {{$ctxargs := ctxargs .Method}}
    for {
      page, res, err := app.gh.{{.Method.Service | pointer}}.{{.Method.Name}}({{with $ctxargs}}{{$ctxargs}}, {{end}}opt)
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
  Action: func(c *cli.Context) {
    fatalln("Not implemented")
  },
},`))
