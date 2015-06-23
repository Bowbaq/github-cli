package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-repository-events`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-repository-events", "list-repository-events <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-issue-events-for-repository", "list-issue-events-for-repository <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-events-for-repo-network", "list-events-for-repo-network <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-public-events-for-an-organization`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-events-for-organization", "list-events-for-organization <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `public-only`, Usage: ``},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-events-performed-by-user", "list-events-performed-by-user <user>")
				}

				user := c.Args().Get(0)
				publicOnly := c.Bool("public-only")

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `public-only`, Usage: ``},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-events-recieved-by-user", "list-events-recieved-by-user <user>")
				}

				user := c.Args().Get(0)
				publicOnly := c.Bool("public-only")

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/activity/events/#list-events-for-an-organization`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-user-events-for-organization", "list-user-events-for-organization <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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
			Usage: `list-notifications lists all notifications for the authenticated user.`,
			Description: `list-notifications lists all notifications for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#list-your-notifications`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `all`, Usage: ``},
				cli.BoolFlag{Name: `participating`, Usage: ``},
				cli.StringFlag{Name: `since`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				opt := &github.NotificationListOptions{
					All:           c.Bool("all"),
					Participating: c.Bool("participating"),
					Since:         now.MustParse(c.String("since")),
				}

				result, res, err := app.gh.Activity.ListNotifications(opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-repository-notifications",
			Usage: `list-repository-notifications lists all notifications in a given repository for the authenticated user.`,
			Description: `list-repository-notifications lists all notifications in a given repository
   for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `since`, Usage: ``},
				cli.BoolFlag{Name: `all`, Usage: ``},
				cli.BoolFlag{Name: `participating`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-repository-notifications", "list-repository-notifications <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.NotificationListOptions{
					All:           c.Bool("all"),
					Participating: c.Bool("participating"),
					Since:         now.MustParse(c.String("since")),
				}

				result, res, err := app.gh.Activity.ListRepositoryNotifications(owner, repo, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "mark-notifications-read",
			Usage: `mark-notifications-read marks all notifications up to lastRead as read.`,
			Description: `mark-notifications-read marks all notifications up to lastRead as read.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#mark-as-read`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `last-read`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				lastRead := now.MustParse(c.String("last-read"))

				res, err := app.gh.Activity.MarkNotificationsRead(lastRead)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "mark-repository-notifications-read",
			Usage: `mark-repository-notifications-read marks all notifications up to lastRead in the specified repository as read.`,
			Description: `mark-repository-notifications-read marks all notifications up to lastRead in
   the specified repository as read.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read-in-a-repository`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `last-read`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "mark-repository-notifications-read", "mark-repository-notifications-read <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				lastRead := now.MustParse(c.String("last-read"))

				res, err := app.gh.Activity.MarkRepositoryNotificationsRead(owner, repo, lastRead)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "get-thread",
			Usage: `get-thread gets the specified notification thread.`,
			Description: `get-thread gets the specified notification thread.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#view-a-single-thread`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get-thread", "get-thread <id>")
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

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "mark-thread-read", "mark-thread-read <id>")
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

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get-thread-subscription", "get-thread-subscription <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Activity.GetThreadSubscription(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "set-thread-subscription",
			Usage: `set-thread-subscription sets the subscription for the specified thread for the authenticated user.`,
			Description: `set-thread-subscription sets the subscription for the specified thread for the
   authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `subscribed`, Usage: ``},
				cli.BoolFlag{Name: `ignored`, Usage: ``},
				cli.StringFlag{Name: `reason`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "set-thread-subscription", "set-thread-subscription <id>")
				}

				id := c.Args().Get(0)
				subscription := &github.Subscription{
					Subscribed: github.Bool(c.Bool("subscribed")),
					Ignored:    github.Bool(c.Bool("ignored")),
					Reason:     github.String(c.String("reason")),
					CreatedAt:  &github.Timestamp{now.MustParse(c.String("created-at"))},
				}

				result, res, err := app.gh.Activity.SetThreadSubscription(id, subscription)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-thread-subscription",
			Usage: `delete-thread-subscription deletes the subscription for the specified thread for the authenticated user.`,
			Description: `delete-thread-subscription deletes the subscription for the specified thread
   for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "delete-thread-subscription", "delete-thread-subscription <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Activity.DeleteThreadSubscription(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-stargazers",
			Usage: `list-stargazers lists people who have starred the specified repo.`,
			Description: `list-stargazers lists people who have starred the specified repo.

   GitHub API Docs: https://developer.github.com/v3/activity/starring/#list-stargazers`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-stargazers", "list-stargazers <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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
			Usage: `list-starred lists all the repos starred by a user.`,
			Description: `list-starred lists all the repos starred by a user.  Passing the empty string
   will list the starred repositories for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/activity/starring/#list-repositories-being-starred`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sort`, Usage: `How to sort the repository list.  Possible values are: created, updated,
pushed, full_name.  Default is "full_name".`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort repositories.  Possible values are: asc, desc.
Default is "asc" when sort is "full_name", otherwise default is "desc".`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-starred", "list-starred <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ActivityListStarredOptions{
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
				}

				var items []github.StarredRepository

				for {
					page, res, err := app.gh.Activity.ListStarred(user, opt)
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
			Name:  "is-starred",
			Usage: `is-starred checks if a repository is starred by authenticated user.`,
			Description: `is-starred checks if a repository is starred by authenticated user.

   GitHub API docs: https://developer.github.com/v3/activity/starring/#check-if-you-are-starring-a-repository`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "is-starred", "is-starred <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/activity/starring/#star-a-repository`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "star", "star <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/activity/starring/#unstar-a-repository`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "unstar", "unstar <owner> <repo>")
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

   GitHub API Docs: http://developer.github.com/v3/activity/watching/#list-watchers`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-watchers", "list-watchers <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#list-repositories-being-watched`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-watched", "list-watched <user>")
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

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#get-a-repository-subscription`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-repository-subscription", "get-repository-subscription <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Activity.GetRepositorySubscription(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "set-repository-subscription",
			Usage: `set-repository-subscription sets the subscription for the specified repository for the authenticated user.`,
			Description: `set-repository-subscription sets the subscription for the specified repository
   for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#set-a-repository-subscription`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `ignored`, Usage: ``},
				cli.StringFlag{Name: `reason`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.BoolFlag{Name: `subscribed`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "set-repository-subscription", "set-repository-subscription <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				subscription := &github.Subscription{
					Reason:     github.String(c.String("reason")),
					CreatedAt:  &github.Timestamp{now.MustParse(c.String("created-at"))},
					Subscribed: github.Bool(c.Bool("subscribed")),
					Ignored:    github.Bool(c.Bool("ignored")),
				}

				result, res, err := app.gh.Activity.SetRepositorySubscription(owner, repo, subscription)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-repository-subscription",
			Usage: `delete-repository-subscription deletes the subscription for the specified repository for the authenticated user.`,
			Description: `delete-repository-subscription deletes the subscription for the specified
   repository for the authenticated user.

   GitHub API Docs: https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "delete-repository-subscription", "delete-repository-subscription <owner> <repo>")
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
