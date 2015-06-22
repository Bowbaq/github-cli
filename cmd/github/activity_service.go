package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var ActivityService = cli.Command{
	Name:     "activity",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list-events",
			Usage: `list-events drinks from the firehose of all public events across GitHub.`,
			Description: `list-events drinks from the firehose of all public events across GitHub.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events
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

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListEvents(opt)
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
			Name:  "list-repository-events",
			Usage: `list-repository-events lists events for a repository.`,
			Description: `list-repository-events lists events for a repository.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-repository-events
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-repository-events <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListRepositoryEvents(owner, repo, opt)
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
			Name:  "list-issue-events-for-repository",
			Usage: `list-issue-events-for-repository lists issue events for a repository.`,
			Description: `list-issue-events-for-repository lists issue events for a repository.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-issue-events-for-repository <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListIssueEventsForRepository(owner, repo, opt)
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
			Name:  "list-events-for-repo-network",
			Usage: `list-events-for-repo-network lists public events for a network of repositories.`,
			Description: `list-events-for-repo-network lists public events for a network of repositories.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-events-for-repo-network <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListEventsForRepoNetwork(owner, repo, opt)
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
			Name:  "list-events-for-organization",
			Usage: `list-events-for-organization lists public events for an organization.`,
			Description: `list-events-for-organization lists public events for an organization.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-an-organization
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-events-for-organization <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListEventsForOrganization(org, opt)
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
			Name:  "list-events-performed-by-user",
			Usage: `list-events-performed-by-user lists the events performed by a user.`,
			Description: `list-events-performed-by-user lists the events performed by a user. If publicOnly is
   true, only public events will be returned.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "public-only"},
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-events-performed-by-user <user>")
				}

				user := c.Args().Get(0)
				publicOnly := c.Bool("public-only")

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListEventsPerformedByUser(user, publicOnly, opt)
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
			Name:  "list-events-recieved-by-user",
			Usage: `list-events-recieved-by-user lists the events recieved by a user.`,
			Description: `list-events-recieved-by-user lists the events recieved by a user. If publicOnly is
   true, only public events will be returned.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "public-only"},
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-events-recieved-by-user <user>")
				}

				user := c.Args().Get(0)
				publicOnly := c.Bool("public-only")

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListEventsRecievedByUser(user, publicOnly, opt)
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
			Name:  "list-user-events-for-organization",
			Usage: `list-user-events-for-organization provides the user’s organization dashboard.`,
			Description: `list-user-events-for-organization provides the user’s organization dashboard. You
   must be authenticated as the user to view this.

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-for-an-organization
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-user-events-for-organization <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Event

				for {
					page, res, err := app.gh.Activity.ListUserEventsForOrganization(org, user, opt)
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
			Name:  "list-notifications",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-repository-notifications",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "mark-notifications-read",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "mark-repository-notifications-read",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-thread",
			Usage: `get-thread gets the specified notification thread.`,
			Description: `get-thread gets the specified notification thread.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#view-a-single-thread
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get-thread <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Activity.GetThread(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "mark-thread-read",
			Usage: `mark-thread-read marks the specified thread as read.`,
			Description: `mark-thread-read marks the specified thread as read.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "mark-thread-read <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Activity.MarkThreadRead(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "get-thread-subscription",
			Usage: `get-thread-subscription checks to see if the authenticated user is subscribed to a thread.`,
			Description: `get-thread-subscription checks to see if the authenticated user is subscribed
   to a thread.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get-thread-subscription <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Activity.GetThreadSubscription(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "set-thread-subscription",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-thread-subscription",
			Usage: `delete-thread-subscription deletes the subscription for the specified thread for the authenticated user.`,
			Description: `delete-thread-subscription deletes the subscription for the specified thread
   for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "delete-thread-subscription <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Activity.DeleteThreadSubscription(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-stargazers",
			Usage: `list-stargazers lists people who have starred the specified repo.`,
			Description: `list-stargazers lists people who have starred the specified repo.

   GitHub API Docs: https://developer.github.com/v3/activity/starring/#list-stargazers
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-stargazers <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Activity.ListStargazers(owner, repo, opt)
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
			Name:  "list-starred",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "is-starred",
			Usage: `is-starred checks if a repository is starred by authenticated user.`,
			Description: `is-starred checks if a repository is starred by authenticated user.

   GitHub API docs: https://developer.github.com/v3/activity/starring/#check-if-you-are-starring-a-repository
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "is-starred <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Activity.IsStarred(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "star",
			Usage: `star a repository as the authenticated user.`,
			Description: `star a repository as the authenticated user.

   GitHub API docs: https://developer.github.com/v3/activity/starring/#star-a-repository
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "star <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				res, err := app.gh.Activity.Star(owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unstar",
			Usage: `unstar a repository as the authenticated user.`,
			Description: `unstar a repository as the authenticated user.

   GitHub API docs: https://developer.github.com/v3/activity/starring/#unstar-a-repository
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "unstar <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				res, err := app.gh.Activity.Unstar(owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-watchers",
			Usage: `list-watchers lists watchers of a particular repo.`,
			Description: `list-watchers lists watchers of a particular repo.

   GitHub API Docs: http://developer.github.com/v3/activity/watching/#list-watchers
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-watchers <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Activity.ListWatchers(owner, repo, opt)
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
			Name:  "list-watched",
			Usage: `list-watched lists the repositories the specified user is watching.`,
			Description: `list-watched lists the repositories the specified user is watching.  Passing
   the empty string will fetch watched repos for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#list-repositories-being-watched
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-watched <user>")
				}

				user := c.Args().Get(0)

				result, res, err := app.gh.Activity.ListWatched(user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-repository-subscription",
			Usage: `get-repository-subscription returns the subscription for the specified repository for the authenticated user.`,
			Description: `get-repository-subscription returns the subscription for the specified
   repository for the authenticated user.  If the authenticated user is not
   watching the repository, a nil Subscription is returned.

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#get-a-repository-subscription
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-repository-subscription <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Activity.GetRepositorySubscription(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "set-repository-subscription",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-repository-subscription",
			Usage: `delete-repository-subscription deletes the subscription for the specified repository for the authenticated user.`,
			Description: `delete-repository-subscription deletes the subscription for the specified
   repository for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "delete-repository-subscription <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				res, err := app.gh.Activity.DeleteRepositorySubscription(owner, repo)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, ActivityService)
}
