package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
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

   GitHub API docs: http://developer.github.com/v3/users/#get-a-single-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get", "get <user>")
				}

				user := c.Args().Get(0)

				result, res, err := app.gh.Users.Get(user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit the authenticated user.`,
			Description: `edit the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/#update-the-authenticated-user`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `hireable`, Usage: ``},
				cli.StringFlag{Name: `company`, Usage: ``},
				cli.StringFlag{Name: `email`, Usage: ``},
				cli.IntFlag{Name: `following`, Usage: ``},
				cli.StringFlag{Name: `type`, Usage: ``},
				cli.StringFlag{Name: `login`, Usage: ``},
				cli.IntFlag{Name: `public-gists`, Usage: ``},
				cli.IntFlag{Name: `followers`, Usage: ``},
				cli.IntFlag{Name: `private-gists`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `blog`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.IntFlag{Name: `total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `collaborators`, Usage: ``},
				cli.StringFlag{Name: `plan-name`, Usage: ``},
				cli.IntFlag{Name: `plan-space`, Usage: ``},
				cli.IntFlag{Name: `plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `gravatar-id`, Usage: ``},
				cli.IntFlag{Name: `public-repos`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.BoolFlag{Name: `site-admin`, Usage: ``},
				cli.StringFlag{Name: `bio`, Usage: ``},
				cli.StringFlag{Name: `location`, Usage: ``},
				cli.IntFlag{Name: `disk-usage`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				user := &github.User{
					Location:          github.String(c.String("location")),
					DiskUsage:         github.Int(c.Int("disk-usage")),
					Hireable:          github.Bool(c.Bool("hireable")),
					Following:         github.Int(c.Int("following")),
					Type:              github.String(c.String("type")),
					Login:             github.String(c.String("login")),
					Company:           github.String(c.String("company")),
					Email:             github.String(c.String("email")),
					Name:              github.String(c.String("name")),
					PublicGists:       github.Int(c.Int("public-gists")),
					Followers:         github.Int(c.Int("followers")),
					PrivateGists:      github.Int(c.Int("private-gists")),
					OwnedPrivateRepos: github.Int(c.Int("owned-private-repos")),
					Collaborators:     github.Int(c.Int("collaborators")),
					GravatarID:        github.String(c.String("gravatar-id")),
					Blog:              github.String(c.String("blog")),
					CreatedAt:         &github.Timestamp{now.MustParse(c.String("created-at"))},
					TotalPrivateRepos: github.Int(c.Int("total-private-repos")),
					ID:                github.Int(c.Int("id")),
					PublicRepos:       github.Int(c.Int("public-repos")),
					UpdatedAt:         &github.Timestamp{now.MustParse(c.String("updated-at"))},
					Bio:               github.String(c.String("bio")),
					SiteAdmin:         github.Bool(c.Bool("site-admin")),
				}

				result, res, err := app.gh.Users.Edit(user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-all",
			Usage: `list-all lists all GitHub users.`,
			Description: `list-all lists all GitHub users.

   GitHub API docs: http://developer.github.com/v3/users/#get-all-users`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `since`, Usage: `ID of the last user seen`},
			},
			Action: func(c *cli.Context) {
				opt := &github.UserListOptions{
					Since: c.Int("since"),
				}

				result, res, err := app.gh.Users.ListAll(opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "promote-site-admin",
			Usage: `promote-site-admin promotes a user to a site administrator of a GitHub Enterprise instance.`,
			Description: `promote-site-admin promotes a user to a site administrator of a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#promote-an-ordinary-user-to-a-site-administrator`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "promote-site-admin", "promote-site-admin <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.PromoteSiteAdmin(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "demote-site-admin",
			Usage: `demote-site-admin demotes a user from site administrator of a GitHub Enterprise instance.`,
			Description: `demote-site-admin demotes a user from site administrator of a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#demote-a-site-administrator-to-an-ordinary-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "demote-site-admin", "demote-site-admin <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.DemoteSiteAdmin(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "suspend",
			Usage: `suspend a user on a GitHub Enterprise instance.`,
			Description: `suspend a user on a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#suspend-a-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "suspend", "suspend <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Suspend(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unsuspend",
			Usage: `unsuspend a user on a GitHub Enterprise instance.`,
			Description: `unsuspend a user on a GitHub Enterprise instance.

   GitHub API docs: https://developer.github.com/v3/users/administration/#unsuspend-a-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "unsuspend", "unsuspend <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Unsuspend(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-emails",
			Usage: `list-emails lists all email addresses for the authenticated user.`,
			Description: `list-emails lists all email addresses for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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
			Usage: `add-emails adds email addresses of the authenticated user.`,
			Description: `add-emails adds email addresses of the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/emails/#add-email-addresses`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `emails`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				emails := c.StringSlice("emails")

				result, res, err := app.gh.Users.AddEmails(emails)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-emails",
			Usage: `delete-emails deletes email addresses from authenticated user.`,
			Description: `delete-emails deletes email addresses from authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/emails/#delete-email-addresses`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `emails`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				emails := c.StringSlice("emails")

				res, err := app.gh.Users.DeleteEmails(emails)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-followers",
			Usage: `list-followers lists the followers for a user.`,
			Description: `list-followers lists the followers for a user.  Passing the empty string will
   fetch followers for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/followers/#list-followers-of-a-user`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-followers", "list-followers <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/users/followers/#list-users-followed-by-another-user`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-following", "list-following <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "is-following", "is-following <user> <target>")
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

   GitHub API docs: http://developer.github.com/v3/users/followers/#follow-a-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "follow", "follow <user>")
				}

				user := c.Args().Get(0)

				res, err := app.gh.Users.Follow(user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unfollow",
			Usage: `unfollow will cause the authenticated user to unfollow the specified user.`,
			Description: `unfollow will cause the authenticated user to unfollow the specified user.

   GitHub API docs: http://developer.github.com/v3/users/followers/#unfollow-a-user`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "unfollow", "unfollow <user>")
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

   GitHub API docs: http://developer.github.com/v3/users/keys/#list-public-keys-for-a-user`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-keys", "list-keys <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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

   GitHub API docs: http://developer.github.com/v3/users/keys/#get-a-single-public-key`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get-key", "get-key <id>")
				}

				id, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				result, res, err := app.gh.Users.GetKey(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-key",
			Usage: `create-key adds a public key for the authenticated user.`,
			Description: `create-key adds a public key for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/users/keys/#create-a-public-key`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `key`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				key := &github.Key{
					ID:    github.Int(c.Int("id")),
					Key:   github.String(c.String("key")),
					Title: github.String(c.String("title")),
				}

				result, res, err := app.gh.Users.CreateKey(key)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-key",
			Usage: `delete-key deletes a public key.`,
			Description: `delete-key deletes a public key.

   GitHub API docs: http://developer.github.com/v3/users/keys/#delete-a-public-key`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "delete-key", "delete-key <id>")
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
