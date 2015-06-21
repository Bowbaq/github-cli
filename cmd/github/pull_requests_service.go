package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var PullRequestsService = cli.Command{
	Name:     "pull-requests",
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
		}, cli.Command{
			Name: "create",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-commits",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-files",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "is-merged",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "merge",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-comments",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, PullRequestsService)
}
