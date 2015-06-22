package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var UsersService = cli.Command{
	Name:     "users",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "get",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-all",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "promote-site-admin",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "demote-site-admin",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "suspend",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "unsuspend",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-emails",
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
				var items []github.UserEmail

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Users.ListEmails(opt)
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
			Name: "add-emails",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-emails",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-followers",
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
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + " list-followers <user>")
				}

				var items []github.User

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Users.ListFollowers(c.Args().Get(0), opt)
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
			Name: "list-following",
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
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + " list-following <user>")
				}

				var items []github.User

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Users.ListFollowing(c.Args().Get(0), opt)
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
			Name: "is-following",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "follow",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "unfollow",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-keys",
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
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + " list-keys <user>")
				}

				var items []github.Key

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Users.ListKeys(c.Args().Get(0), opt)
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
			Name: "get-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, UsersService)
}
