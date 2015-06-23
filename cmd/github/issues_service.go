package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
	"github.com/kr/pretty"
)

var IssuesService = cli.Command{
	Name:     "issues",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list the issues for the authenticated user.`,
			Description: `list the issues for the authenticated user.  If all is true, list issues
   across all the user's visible repositories including owned, member, and
   organization repositories; if false, list only owned and member
   repositories.

   GitHub API docs: http://developer.github.com/v3/issues/#list-issues`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `all`, Usage: ``},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort issues.  Possible values are: asc, desc.
Default is "asc".`},
				cli.StringFlag{Name: `since`, Usage: `Since filters issues by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `filter`, Usage: `Filter specifies which issues to list.  Possible values are: assigned,
created, mentioned, subscribed, all.  Default is "assigned".`},
				cli.StringFlag{Name: `state`, Usage: `State filters issues based on their state.  Possible values are: open,
closed.  Default is "open".`},
				cli.StringSliceFlag{Name: `labels`, Usage: `Labels filters issues based on their label.`},
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort issues.  Possible values are: created, updated,
and comments.  Default value is "created".`},
			},
			Action: func(c *cli.Context) {
				all := c.Bool("all")

				opt := &github.IssueListOptions{
					Filter:    c.String("filter"),
					State:     c.String("state"),
					Labels:    c.StringSlice("labels"),
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
					Since:     now.MustParse(c.String("since")),
				}

				var items []github.Issue

				for {
					page, res, err := app.gh.Issues.List(all, opt)
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
			Name:  "list-by-org",
			Usage: `list-by-org fetches the issues in the specified organization for the authenticated user.`,
			Description: `list-by-org fetches the issues in the specified organization for the
   authenticated user.

   GitHub API docs: http://developer.github.com/v3/issues/#list-issues`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort issues.  Possible values are: created, updated,
and comments.  Default value is "created".`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort issues.  Possible values are: asc, desc.
Default is "asc".`},
				cli.StringFlag{Name: `since`, Usage: `Since filters issues by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `filter`, Usage: `Filter specifies which issues to list.  Possible values are: assigned,
created, mentioned, subscribed, all.  Default is "assigned".`},
				cli.StringFlag{Name: `state`, Usage: `State filters issues based on their state.  Possible values are: open,
closed.  Default is "open".`},
				cli.StringSliceFlag{Name: `labels`, Usage: `Labels filters issues based on their label.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-by-org", "list-by-org <org>")
				}

				org := c.Args().Get(0)
				opt := &github.IssueListOptions{
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
					Since:     now.MustParse(c.String("since")),
					Filter:    c.String("filter"),
					State:     c.String("state"),
					Labels:    c.StringSlice("labels"),
				}

				var items []github.Issue

				for {
					page, res, err := app.gh.Issues.ListByOrg(org, opt)
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
			Name:  "list-by-repo",
			Usage: `list-by-repo lists the issues for the specified repository.`,
			Description: `list-by-repo lists the issues for the specified repository.

   GitHub API docs: http://developer.github.com/v3/issues/#list-issues-for-a-repository`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `state`, Usage: `State filters issues based on their state.  Possible values are: open,
closed.  Default is "open".`},
				cli.StringFlag{Name: `mentioned`, Usage: `Assignee filters issues to those mentioned a specific user.`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort issues.  Possible values are: asc, desc.
Default is "asc".`},
				cli.StringFlag{Name: `since`, Usage: `Since filters issues by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `milestone`, Usage: `Milestone limits issues for the specified milestone.  Possible values are
a milestone number, "none" for issues with no milestone, "*" for issues
with any milestone.`},
				cli.StringFlag{Name: `assignee`, Usage: `Assignee filters issues based on their assignee.  Possible values are a
user name, "none" for issues that are not assigned, "*" for issues with
any assigned user.`},
				cli.StringFlag{Name: `creator`, Usage: `Assignee filters issues based on their creator.`},
				cli.StringSliceFlag{Name: `labels`, Usage: `Labels filters issues based on their label.`},
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort issues.  Possible values are: created, updated,
and comments.  Default value is "created".`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-by-repo", "list-by-repo <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.IssueListByRepoOptions{
					Creator:   c.String("creator"),
					Labels:    c.StringSlice("labels"),
					Sort:      c.String("sort"),
					Milestone: c.String("milestone"),
					Assignee:  c.String("assignee"),
					Direction: c.String("direction"),
					Since:     now.MustParse(c.String("since")),
					State:     c.String("state"),
					Mentioned: c.String("mentioned"),
				}

				var items []github.Issue

				for {
					page, res, err := app.gh.Issues.ListByRepo(owner, repo, opt)
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
			Name:  "get",
			Usage: `get a single issue.`,
			Description: `get a single issue.

   GitHub API docs: http://developer.github.com/v3/issues/#get-a-single-issue`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get", "get <owner> <repo> <number>")
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
			Usage: `create a new issue on the specified repository.`,
			Description: `create a new issue on the specified repository.

   GitHub API docs: http://developer.github.com/v3/issues/#create-an-issue`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringSliceFlag{Name: `labels`, Usage: ``},
				cli.StringFlag{Name: `assignee`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.IntFlag{Name: `milestone`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create", "create <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				issue := &github.IssueRequest{
					Title:     github.String(c.String("title")),
					Body:      github.String(c.String("body")),
					Labels:    stringSlicePointer(c.StringSlice("labels")),
					Assignee:  github.String(c.String("assignee")),
					State:     github.String(c.String("state")),
					Milestone: github.Int(c.Int("milestone")),
				}

				result, res, err := app.gh.Issues.Create(owner, repo, issue)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit an issue.`,
			Description: `edit an issue.

   GitHub API docs: http://developer.github.com/v3/issues/#edit-an-issue`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringSliceFlag{Name: `labels`, Usage: ``},
				cli.StringFlag{Name: `assignee`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.IntFlag{Name: `milestone`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit", "edit <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				issue := &github.IssueRequest{
					Milestone: github.Int(c.Int("milestone")),
					Title:     github.String(c.String("title")),
					Body:      github.String(c.String("body")),
					Labels:    stringSlicePointer(c.StringSlice("labels")),
					Assignee:  github.String(c.String("assignee")),
					State:     github.String(c.String("state")),
				}

				result, res, err := app.gh.Issues.Edit(owner, repo, number, issue)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-assignees",
			Usage: `list-assignees fetches all available assignees (owners and collaborators) to which issues may be assigned.`,
			Description: `list-assignees fetches all available assignees (owners and collaborators) to
   which issues may be assigned.

   GitHub API docs: http://developer.github.com/v3/issues/assignees/#list-assignees`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-assignees", "list-assignees <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/issues/assignees/#check-assignee`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "is-assignee", "is-assignee <owner> <repo> <user>")
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
			Usage: `list-comments lists all comments on the specified issue.`,
			Description: `list-comments lists all comments on the specified issue.  Specifying an issue
   number of 0 will return all comments on all issues for the repository.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#list-comments-on-an-issue`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort comments.  Possible values are: asc, desc.`},
				cli.StringFlag{Name: `since`, Usage: `Since filters comments by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort comments.  Possible values are: created, updated.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-comments", "list-comments <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.IssueListCommentsOptions{
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
					Since:     now.MustParse(c.String("since")),
				}

				var items []github.IssueComment

				for {
					page, res, err := app.gh.Issues.ListComments(owner, repo, number, opt)
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
			Name:  "get-comment",
			Usage: `get-comment fetches the specified issue comment.`,
			Description: `get-comment fetches the specified issue comment.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#get-a-single-comment`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-comment", "get-comment <owner> <repo> <id>")
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
			Usage: `create-comment creates a new comment on the specified issue.`,
			Description: `create-comment creates a new comment on the specified issue.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#create-a-comment`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-comment", "create-comment <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				comment := &github.IssueComment{
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
				}

				result, res, err := app.gh.Issues.CreateComment(owner, repo, number, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-comment",
			Usage: `edit-comment updates an issue comment.`,
			Description: `edit-comment updates an issue comment.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#edit-a-comment`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-comment", "edit-comment <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				comment := &github.IssueComment{
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Issues.EditComment(owner, repo, id, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes an issue comment.`,
			Description: `delete-comment deletes an issue comment.

   GitHub API docs: http://developer.github.com/v3/issues/comments/#delete-a-comment`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-comment", "delete-comment <owner> <repo> <id>")
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

   GitHub API docs: https://developer.github.com/v3/issues/events/#list-events-for-an-issue`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-issue-events", "list-issue-events <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: https://developer.github.com/v3/issues/events/#list-events-for-a-repository`,
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

   GitHub API docs: https://developer.github.com/v3/issues/events/#get-a-single-event`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-event", "get-event <owner> <repo> <id>")
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

   GitHub API docs: http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-labels", "list-labels <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/issues/labels/#get-a-single-label`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-label", "get-label <owner> <repo> <name>")
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
			Usage: `create-label creates a new label on the specified repository.`,
			Description: `create-label creates a new label on the specified repository.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#create-a-label`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `color`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-label", "create-label <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				label := &github.Label{
					Name:  github.String(c.String("name")),
					Color: github.String(c.String("color")),
				}

				result, res, err := app.gh.Issues.CreateLabel(owner, repo, label)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-label",
			Usage: `edit-label edits a label.`,
			Description: `edit-label edits a label.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#update-a-label`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `color`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-label", "edit-label <owner> <repo> <name>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				name := c.Args().Get(2)
				label := &github.Label{
					Name:  github.String(c.String("name")),
					Color: github.String(c.String("color")),
				}

				result, res, err := app.gh.Issues.EditLabel(owner, repo, name, label)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-label",
			Usage: `delete-label deletes a label.`,
			Description: `delete-label deletes a label.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#delete-a-label`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-label", "delete-label <owner> <repo> <name>")
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

   GitHub API docs: http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-labels-by-issue", "list-labels-by-issue <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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
			Usage: `add-labels-to-issue adds labels to an issue.`,
			Description: `add-labels-to-issue adds labels to an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `labels`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "add-labels-to-issue", "add-labels-to-issue <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				labels := c.StringSlice("labels")

				result, res, err := app.gh.Issues.AddLabelsToIssue(owner, repo, number, labels)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "remove-label-for-issue",
			Usage: `remove-label-for-issue removes a label for an issue.`,
			Description: `remove-label-for-issue removes a label for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					showHelp(c, "remove-label-for-issue", "remove-label-for-issue <owner> <repo> <number> <label>")
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
			Usage: `replace-labels-for-issue replaces all labels for an issue.`,
			Description: `replace-labels-for-issue replaces all labels for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `labels`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "replace-labels-for-issue", "replace-labels-for-issue <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				labels := c.StringSlice("labels")

				result, res, err := app.gh.Issues.ReplaceLabelsForIssue(owner, repo, number, labels)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "remove-labels-for-issue",
			Usage: `remove-labels-for-issue removes all labels for an issue.`,
			Description: `remove-labels-for-issue removes all labels for an issue.

   GitHub API docs: http://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "remove-labels-for-issue", "remove-labels-for-issue <owner> <repo> <number>")
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

   GitHub API docs: http://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-labels-for-milestone", "list-labels-for-milestone <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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
			Usage: `list-milestones lists all milestones for a repository.`,
			Description: `list-milestones lists all milestones for a repository.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `state`, Usage: `State filters milestones based on their state. Possible values are:
open, closed. Default is "open".`},
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort milestones. Possible values are: due_date, completeness.
Default value is "due_date".`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort milestones. Possible values are: asc, desc.
Default is "asc".`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-milestones", "list-milestones <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.MilestoneListOptions{
					Direction: c.String("direction"),
					State:     c.String("state"),
					Sort:      c.String("sort"),
				}

				result, res, err := app.gh.Issues.ListMilestones(owner, repo, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-milestone",
			Usage: `get-milestone gets a single milestone.`,
			Description: `get-milestone gets a single milestone.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#get-a-single-milestone`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-milestone", "get-milestone <owner> <repo> <number>")
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
			Usage: `create-milestone creates a new milestone on the specified repository.`,
			Description: `create-milestone creates a new milestone on the specified repository.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#create-a-milestone`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `due-on`, Usage: ``},
				cli.IntFlag{Name: `number`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.IntFlag{Name: `creator-private-gists`, Usage: ``},
				cli.StringFlag{Name: `creator-name`, Usage: ``},
				cli.IntFlag{Name: `creator-public-gists`, Usage: ``},
				cli.IntFlag{Name: `creator-followers`, Usage: ``},
				cli.IntFlag{Name: `creator-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-collaborators`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `creator-plan-name`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-space`, Usage: ``},
				cli.StringFlag{Name: `creator-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `creator-blog`, Usage: ``},
				cli.StringFlag{Name: `creator-created-at`, Usage: ``},
				cli.IntFlag{Name: `creator-id`, Usage: ``},
				cli.IntFlag{Name: `creator-public-repos`, Usage: ``},
				cli.StringFlag{Name: `creator-updated-at`, Usage: ``},
				cli.StringFlag{Name: `creator-bio`, Usage: ``},
				cli.BoolFlag{Name: `creator-site-admin`, Usage: ``},
				cli.StringFlag{Name: `creator-location`, Usage: ``},
				cli.IntFlag{Name: `creator-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `creator-hireable`, Usage: ``},
				cli.StringFlag{Name: `creator-email`, Usage: ``},
				cli.IntFlag{Name: `creator-following`, Usage: ``},
				cli.StringFlag{Name: `creator-type`, Usage: ``},
				cli.StringFlag{Name: `creator-login`, Usage: ``},
				cli.StringFlag{Name: `creator-company`, Usage: ``},
				cli.IntFlag{Name: `closed-issues`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.IntFlag{Name: `open-issues`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-milestone", "create-milestone <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				milestone := &github.Milestone{
					DueOn:        timePointer(now.MustParse(c.String("due-on"))),
					Number:       github.Int(c.Int("number")),
					State:        github.String(c.String("state")),
					ClosedIssues: github.Int(c.Int("closed-issues")),
					CreatedAt:    timePointer(now.MustParse(c.String("created-at"))),
					Title:        github.String(c.String("title")),
					Description:  github.String(c.String("description")),
					OpenIssues:   github.Int(c.Int("open-issues")),
					UpdatedAt:    timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Issues.CreateMilestone(owner, repo, milestone)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-milestone",
			Usage: `edit-milestone edits a milestone.`,
			Description: `edit-milestone edits a milestone.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#update-a-milestone`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `closed-issues`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `due-on`, Usage: ``},
				cli.IntFlag{Name: `number`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.StringFlag{Name: `creator-location`, Usage: ``},
				cli.IntFlag{Name: `creator-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `creator-hireable`, Usage: ``},
				cli.StringFlag{Name: `creator-company`, Usage: ``},
				cli.StringFlag{Name: `creator-email`, Usage: ``},
				cli.IntFlag{Name: `creator-following`, Usage: ``},
				cli.StringFlag{Name: `creator-type`, Usage: ``},
				cli.StringFlag{Name: `creator-login`, Usage: ``},
				cli.IntFlag{Name: `creator-public-gists`, Usage: ``},
				cli.IntFlag{Name: `creator-followers`, Usage: ``},
				cli.IntFlag{Name: `creator-private-gists`, Usage: ``},
				cli.StringFlag{Name: `creator-name`, Usage: ``},
				cli.StringFlag{Name: `creator-blog`, Usage: ``},
				cli.StringFlag{Name: `creator-created-at`, Usage: ``},
				cli.IntFlag{Name: `creator-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-collaborators`, Usage: ``},
				cli.StringFlag{Name: `creator-plan-name`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-space`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `creator-gravatar-id`, Usage: ``},
				cli.IntFlag{Name: `creator-public-repos`, Usage: ``},
				cli.StringFlag{Name: `creator-updated-at`, Usage: ``},
				cli.IntFlag{Name: `creator-id`, Usage: ``},
				cli.BoolFlag{Name: `creator-site-admin`, Usage: ``},
				cli.StringFlag{Name: `creator-bio`, Usage: ``},
				cli.IntFlag{Name: `open-issues`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-milestone", "edit-milestone <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				milestone := &github.Milestone{
					ClosedIssues: github.Int(c.Int("closed-issues")),
					CreatedAt:    timePointer(now.MustParse(c.String("created-at"))),
					DueOn:        timePointer(now.MustParse(c.String("due-on"))),
					Number:       github.Int(c.Int("number")),
					State:        github.String(c.String("state")),
					OpenIssues:   github.Int(c.Int("open-issues")),
					UpdatedAt:    timePointer(now.MustParse(c.String("updated-at"))),
					Title:        github.String(c.String("title")),
					Description:  github.String(c.String("description")),
				}

				result, res, err := app.gh.Issues.EditMilestone(owner, repo, number, milestone)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-milestone",
			Usage: `delete-milestone deletes a milestone.`,
			Description: `delete-milestone deletes a milestone.

   GitHub API docs: https://developer.github.com/v3/issues/milestones/#delete-a-milestone`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-milestone", "delete-milestone <owner> <repo> <number>")
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
