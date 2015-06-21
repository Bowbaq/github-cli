package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var IssuesService = cli.Command{
	Name:     "issues",
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
			Name: "list-by-org",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-issues",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-by-repo",
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
			Name: "list-assignees",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "is-assignee",
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
		}, cli.Command{
			Name: "list-issue-events",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-repository-events",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-event",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-labels",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-label",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-label",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-label",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-label",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-labels-by-issue",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "add-labels-to-issue",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "remove-label-for-issue",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "replace-labels-for-issue",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "remove-labels-for-issue",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-labels-for-milestone",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-milestones",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-milestone",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-milestone",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-milestone",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-milestone",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, IssuesService)
}
