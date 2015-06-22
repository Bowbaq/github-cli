package main

import (
	"fmt"
	"strconv"

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
			Name:  "list",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get",
			Usage: `get a single pull request.`,
			Description: `get a single pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#get-a-single-pull-request
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get", "get <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.PullRequests.Get(owner, repo, number)
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
			Name:  "list-commits",
			Usage: `list-commits lists the commits in a pull request.`,
			Description: `list-commits lists the commits in a pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-commits", "list-commits <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepositoryCommit

				for {
					page, res, err := app.gh.PullRequests.ListCommits(owner, repo, number, opt)
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
			Name:  "list-files",
			Usage: `list-files lists the files in a pull request.`,
			Description: `list-files lists the files in a pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#list-pull-requests-files
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-files", "list-files <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.CommitFile

				for {
					page, res, err := app.gh.PullRequests.ListFiles(owner, repo, number, opt)
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
			Name:  "is-merged",
			Usage: `is-merged checks if a pull request has been merged.`,
			Description: `is-merged checks if a pull request has been merged.

   GitHub API docs: https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "is-merged", "is-merged <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.PullRequests.IsMerged(owner, repo, number)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "merge",
			Usage: `merge a pull request (merge Button™).`,
			Description: `merge a pull request (merge Button™).

   GitHub API docs: https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-buttontrade
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					showHelp(c, "merge", "merge <owner> <repo> <number> <commit-message>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				commitMessage := c.Args().Get(3)

				result, res, err := app.gh.PullRequests.Merge(owner, repo, number, commitMessage)
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
			Usage: `get-comment fetches the specified pull request comment.`,
			Description: `get-comment fetches the specified pull request comment.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#get-a-single-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-comment", "get-comment <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.PullRequests.GetComment(owner, repo, number)
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
			Usage: `delete-comment deletes a pull request comment.`,
			Description: `delete-comment deletes a pull request comment.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#delete-a-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-comment", "delete-comment <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.PullRequests.DeleteComment(owner, repo, number)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, PullRequestsService)
}
