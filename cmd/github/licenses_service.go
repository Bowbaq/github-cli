package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/kr/pretty"
)

var LicensesService = cli.Command{
	Name:     "licenses",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list popular open source licenses.`,
			Description: `list popular open source licenses.

   GitHub API docs: https://developer.github.com/v3/licenses/#list-all-licenses
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {

				result, res, err := app.gh.Licenses.List()
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get",
			Usage: `Fetch extended metadata for one license.`,
			Description: `Fetch extended metadata for one license.

   GitHub API docs: https://developer.github.com/v3/licenses/#get-an-individual-license
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get", "get <license-name>")
				}

				licenseName := c.Args().Get(0)

				result, res, err := app.gh.Licenses.Get(licenseName)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, LicensesService)
}
