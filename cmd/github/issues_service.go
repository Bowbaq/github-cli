package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var IssuesService = cli.Command{
	Name:     "issues",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-by-org",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-by-repo",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get",
			Usage: `get a single issue.`,
			Description: `get a single issue.

   GitHub API docs: http://developer.github.com/v3/issues/#get-a-single-issue
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Issues.Get(owner, repo, number)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-assignees",
			Usage: `list-assignees fetches all available assignees (owners and collaborators) to which issues may be assigned.`,
			Description: `list-assignees fetches all available assignees (owners and collaborators) to
   which issues may be assigned.

   GitHub API docs: http://developer.github.com/v3/issues/assignees/#list-assignees
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-assignees <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Issues.ListAssignees(owner, repo, opt)
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
			Name:  "is-assignee",
			Usage: `is-assignee checks if a user is an assignee for the specified repository.`,
			Description: `is-assignee checks if a user is an assignee for the specified repository.

   GitHub API docs: http://developer.github.com/v3/issues/assignees/#check-assignee
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "is-assignee <owner> <repo> <user>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				user := c.Args().Get(2)

				result, res, err := app.gh.Issues.IsAssignee(owner, repo, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-comments",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-comment",
			Usage: `get-comment fetches the specified issue comment.`,
			Description: `get-comment fetches the specified issue comment.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#get-a-single-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-comment <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Issues.GetComment(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-comment",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-comment",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes an issue comment.`,
			Description: `delete-comment deletes an issue comment.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#delete-a-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-comment <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Issues.DeleteComment(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-issue-events",
			Usage: `list-issue-events lists events for the specified issue.`,
			Description: `list-issue-events lists events for the specified issue.

   GitHub API docs: https://developer.github.com/v3/issues/events/#list-events-for-an-issue
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-issue-events <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.IssueEvent

				for {
					page, res, err := app.gh.Issues.ListIssueEvents(owner, repo, number, opt)
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
			Usage: `list-repository-events lists events for the specified repository.`,
			Description: `list-repository-events lists events for the specified repository.

   GitHub API docs: https://developer.github.com/v3/issues/events/#list-events-for-a-repository
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

				var items []github.IssueEvent

				for {
					page, res, err := app.gh.Issues.ListRepositoryEvents(owner, repo, opt)
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
			Name:  "get-event",
			Usage: `get-event returns the specified issue event.`,
			Description: `get-event returns the specified issue event.

   GitHub API docs: https://developer.github.com/v3/issues/events/#get-a-single-event
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-event <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Issues.GetEvent(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-labels",
			Usage: `list-labels lists all labels for a repository.`,
			Description: `list-labels lists all labels for a repository.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-labels <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Label

				for {
					page, res, err := app.gh.Issues.ListLabels(owner, repo, opt)
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
			Name:  "get-label",
			Usage: `get-label gets a single label.`,
			Description: `get-label gets a single label.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#get-a-single-label
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-label <owner> <repo> <name>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				name := c.Args().Get(2)

				result, res, err := app.gh.Issues.GetLabel(owner, repo, name)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-label",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-label",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-label",
			Usage: `delete-label deletes a label.`,
			Description: `delete-label deletes a label.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#delete-a-label
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-label <owner> <repo> <name>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				name := c.Args().Get(2)

				res, err := app.gh.Issues.DeleteLabel(owner, repo, name)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-labels-by-issue",
			Usage: `list-labels-by-issue lists all labels for an issue.`,
			Description: `list-labels-by-issue lists all labels for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-labels-by-issue <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Label

				for {
					page, res, err := app.gh.Issues.ListLabelsByIssue(owner, repo, number, opt)
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
			Name:  "add-labels-to-issue",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "remove-label-for-issue",
			Usage: `remove-label-for-issue removes a label for an issue.`,
			Description: `remove-label-for-issue removes a label for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					fatalln("Usage: " + c.App.Name + "remove-label-for-issue <owner> <repo> <number> <label>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				label := c.Args().Get(3)

				res, err := app.gh.Issues.RemoveLabelForIssue(owner, repo, number, label)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "replace-labels-for-issue",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "remove-labels-for-issue",
			Usage: `remove-labels-for-issue removes all labels for an issue.`,
			Description: `remove-labels-for-issue removes all labels for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "remove-labels-for-issue <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Issues.RemoveLabelsForIssue(owner, repo, number)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-labels-for-milestone",
			Usage: `list-labels-for-milestone lists labels for every issue in a milestone.`,
			Description: `list-labels-for-milestone lists labels for every issue in a milestone.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-labels-for-milestone <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Label

				for {
					page, res, err := app.gh.Issues.ListLabelsForMilestone(owner, repo, number, opt)
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
			Name:  "list-milestones",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-milestone",
			Usage: `get-milestone gets a single milestone.`,
			Description: `get-milestone gets a single milestone.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#get-a-single-milestone
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-milestone <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Issues.GetMilestone(owner, repo, number)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-milestone",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-milestone",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-milestone",
			Usage: `delete-milestone deletes a milestone.`,
			Description: `delete-milestone deletes a milestone.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#delete-a-milestone
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-milestone <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Issues.DeleteMilestone(owner, repo, number)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, IssuesService)
}
