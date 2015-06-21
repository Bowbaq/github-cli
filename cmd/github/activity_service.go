package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

var ActivityService = cli.Command{
	Name:     "activity",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "list-events",
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
				var items []github.Event

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Activity.ListEvents(opt)
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
		}, cli.Command{
			Name: "list-repository-events",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-issue-events-for-repository",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-events-for-repo-network",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-events-for-organization",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-events-performed-by-user",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-events-recieved-by-user",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-user-events-for-organization",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-notifications",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-repository-notifications",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "mark-notifications-read",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "mark-repository-notifications-read",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-thread",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "mark-thread-read",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-thread-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "set-thread-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-thread-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-stargazers",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-starred",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "is-starred",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "star",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "unstar",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-watchers",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-watched",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-repository-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "set-repository-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-repository-subscription",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, ActivityService)
}
