package main

import "github.com/codegangsta/cli"

var SearchService = cli.Command{
	Name:     "search",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "repositories",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "issues",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "users",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "code",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, SearchService)
}
