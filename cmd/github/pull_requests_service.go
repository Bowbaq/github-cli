package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var PullRequestsService = cli.Command{
	Name:     "pull-requests",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "list",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-commits",
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
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-commits <owner> <repo> <number>")
				}

				var items []github.RepositoryCommit

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.PullRequests.ListCommits(c.Args().Get(0), c.Args().Get(1), parseInt(c.Args().Get(2)), opt)
					checkResponse(res.Response, err)

					items = append(items, page...)
					if res.NextPage == 0 || !c.Bool("all") {
						break
					}
					opt.Page = res.NextPage
				}

				fmt.Printf("%# v", pretty.Formatter(items))
			},
		}, cli.Command{
			Name: "list-files",
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
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-files <owner> <repo> <number>")
				}

				var items []github.CommitFile

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.PullRequests.ListFiles(c.Args().Get(0), c.Args().Get(1), parseInt(c.Args().Get(2)), opt)
					checkResponse(res.Response, err)

					items = append(items, page...)
					if res.NextPage == 0 || !c.Bool("all") {
						break
					}
					opt.Page = res.NextPage
				}

				fmt.Printf("%# v", pretty.Formatter(items))
			},
		}, cli.Command{
			Name: "is-merged",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "merge",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-comments",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, PullRequestsService)
}
