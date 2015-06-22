package main

import (
	"fmt"
	"strconv"

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
			Name:  "get",
			Usage: `get fetches a user.`,
			Description: `get fetches a user.  Passing the empty string will fetch the authenticated
   user.

   GitHub API docs: http://developer.github.com/v3/users/#get-a-single-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get <user>")
				}

				user := c.Args().Get(0)

				result, res, err := app.gh.Users.Get(user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-all",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "promote-site-admin",
			Usage: `promote-site-admin promotes a user to a site administrator of a GitHub Enterprise instance.`,
			Description: `promote-site-admin promotes a user to a site administrator of a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#promote-an-ordinary-user-to-a-site-administrator
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "promote-site-admin <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.PromoteSiteAdmin(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "demote-site-admin",
			Usage: `demote-site-admin demotes a user from site administrator of a GitHub Enterprise instance.`,
			Description: `demote-site-admin demotes a user from site administrator of a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#demote-a-site-administrator-to-an-ordinary-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "demote-site-admin <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.DemoteSiteAdmin(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "suspend",
			Usage: `suspend a user on a GitHub Enterprise instance.`,
			Description: `suspend a user on a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#suspend-a-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "suspend <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Suspend(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unsuspend",
			Usage: `unsuspend a user on a GitHub Enterprise instance.`,
			Description: `unsuspend a user on a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#unsuspend-a-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "unsuspend <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Unsuspend(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-emails",
			Usage: `list-emails lists all email addresses for the authenticated user.`,
			Description: `list-emails lists all email addresses for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.UserEmail

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
			Name:  "add-emails",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-emails",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-followers",
			Usage: `list-followers lists the followers for a user.`,
			Description: `list-followers lists the followers for a user.  Passing the empty string will
   fetch followers for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/followers/#list-followers-of-a-user
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-followers <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Users.ListFollowers(user, opt)
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
			Name:  "list-following",
			Usage: `list-following lists the people that a user is following.`,
			Description: `list-following lists the people that a user is following.  Passing the empty
   string will list people the authenticated user is following.

   GitHub API docs: http://developer.github.com/v3/users/followers/#list-users-followed-by-another-user
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-following <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Users.ListFollowing(user, opt)
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
			Name:  "is-following",
			Usage: `is-following checks if "user" is following "target".`,
			Description: `is-following checks if "user" is following "target".  Passing the empty
   string for "user" will check if the authenticated user is following "target".

   GitHub API docs: http://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "is-following <user> <target>")
				}

				user := c.Args().Get(0)
				target := c.Args().Get(1)

				result, res, err := app.gh.Users.IsFollowing(user, target)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "follow",
			Usage: `follow will cause the authenticated user to follow the specified user.`,
			Description: `follow will cause the authenticated user to follow the specified user.

   GitHub API docs: http://developer.github.com/v3/users/followers/#follow-a-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "follow <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Follow(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unfollow",
			Usage: `unfollow will cause the authenticated user to unfollow the specified user.`,
			Description: `unfollow will cause the authenticated user to unfollow the specified user.

   GitHub API docs: http://developer.github.com/v3/users/followers/#unfollow-a-user
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "unfollow <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Unfollow(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-keys",
			Usage: `list-keys lists the verified public keys for a user.`,
			Description: `list-keys lists the verified public keys for a user.  Passing the empty
   string will fetch keys for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/keys/#list-public-keys-for-a-user
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-keys <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Key

				for {
					page, res, err := app.gh.Users.ListKeys(user, opt)
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
			Name:  "get-key",
			Usage: `get-key fetches a single public key.`,
			Description: `get-key fetches a single public key.

   GitHub API docs: http://developer.github.com/v3/users/keys/#get-a-single-public-key
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get-key <id>")
				}

				id, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				result, res, err := app.gh.Users.GetKey(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-key",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-key",
			Usage: `delete-key deletes a public key.`,
			Description: `delete-key deletes a public key.

   GitHub API docs: http://developer.github.com/v3/users/keys/#delete-a-public-key
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "delete-key <id>")
				}

				id, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				res, err := app.gh.Users.DeleteKey(id)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, UsersService)
}
