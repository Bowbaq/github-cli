package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var GistsService = cli.Command{
	Name:     "gists",
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
			Name:  "list-all",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-starred",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get",
			Usage: `get a single gist.`,
			Description: `get a single gist.

   GitHub API docs: http://developer.github.com/v3/gists/#get-a-single-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get", "get <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Gists.Get(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-revision",
			Usage: `Get a specific revision of a gist.`,
			Description: `Get a specific revision of a gist.

   GitHub API docs: https://developer.github.com/v3/gists/#get-a-specific-revision-of-a-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-revision", "get-revision <id> <sha>")
				}

				id := c.Args().Get(0)
				sha := c.Args().Get(1)

				result, res, err := app.gh.Gists.GetRevision(id, sha)
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
			Name:  "delete",
			Usage: `delete a gist.`,
			Description: `delete a gist.

   GitHub API docs: http://developer.github.com/v3/gists/#delete-a-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "delete", "delete <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Gists.Delete(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "star",
			Usage: `star a gist on behalf of authenticated user.`,
			Description: `star a gist on behalf of authenticated user.

   GitHub API docs: http://developer.github.com/v3/gists/#star-a-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "star", "star <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Gists.Star(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "unstar",
			Usage: `unstar a gist on a behalf of authenticated user.`,
			Description: `unstar a gist on a behalf of authenticated user.

   Github API docs: http://developer.github.com/v3/gists/#unstar-a-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "unstar", "unstar <id>")
				}

				id := c.Args().Get(0)

				res, err := app.gh.Gists.Unstar(id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "is-starred",
			Usage: `is-starred checks if a gist is starred by authenticated user.`,
			Description: `is-starred checks if a gist is starred by authenticated user.

   GitHub API docs: http://developer.github.com/v3/gists/#check-if-a-gist-is-starred
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "is-starred", "is-starred <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Gists.IsStarred(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "fork",
			Usage: `fork a gist.`,
			Description: `fork a gist.

   GitHub API docs: http://developer.github.com/v3/gists/#fork-a-gist
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "fork", "fork <id>")
				}

				id := c.Args().Get(0)

				result, res, err := app.gh.Gists.Fork(id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-comments",
			Usage: `list-comments lists all comments for a gist.`,
			Description: `list-comments lists all comments for a gist.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#list-comments-on-a-gist
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-comments", "list-comments <gist-id>")
				}

				gistID := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.GistComment

				for {
					page, res, err := app.gh.Gists.ListComments(gistID, opt)
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
			Usage: `get-comment retrieves a single comment from a gist.`,
			Description: `get-comment retrieves a single comment from a gist.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#get-a-single-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-comment", "get-comment <gist-id> <comment-id>")
				}

				gistID := c.Args().Get(0)
				commentID, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				result, res, err := app.gh.Gists.GetComment(gistID, commentID)
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
			Usage: `delete-comment deletes a gist comment.`,
			Description: `delete-comment deletes a gist comment.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#delete-a-comment
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "delete-comment", "delete-comment <gist-id> <comment-id>")
				}

				gistID := c.Args().Get(0)
				commentID, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				res, err := app.gh.Gists.DeleteComment(gistID, commentID)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, GistsService)
}
