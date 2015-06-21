package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var LicensesService = cli.Command{
	Name:     "licenses",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "list",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, LicensesService)
}
