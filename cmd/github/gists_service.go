package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
	"github.com/kr/pretty"
)

var GistsService = cli.Command{
	Name:     "gists",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list gists for a user.`,
			Description: `list gists for a user. Passing the empty string will list
   all public gists if called anonymously. However, if the call
   is authenticated, it will returns all gists for the authenticated
   user.

   GitHub API docs: http://developer.github.com/v3/gists/#list-gists`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `since`, Usage: `Since filters Gists by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list", "list <user>")
				}

				user := c.Args().Get(0)
				opt := &github.GistListOptions{
					Since: now.MustParse(c.String("since")),
				}

				var items []github.Gist

				for {
					page, res, err := app.gh.Gists.List(user, opt)
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
			Name:  "list-all",
			Usage: `list-all lists all public gists.`,
			Description: `list-all lists all public gists.

   GitHub API docs: http://developer.github.com/v3/gists/#list-gists`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `since`, Usage: `Since filters Gists by time.`},
			},
			Action: func(c *cli.Context) {
				opt := &github.GistListOptions{
					Since: now.MustParse(c.String("since")),
				}

				var items []github.Gist

				for {
					page, res, err := app.gh.Gists.ListAll(opt)
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
			Usage: `list-starred lists starred gists of authenticated user.`,
			Description: `list-starred lists starred gists of authenticated user.

   GitHub API docs: http://developer.github.com/v3/gists/#list-gists`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `since`, Usage: `Since filters Gists by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.GistListOptions{
					Since: now.MustParse(c.String("since")),
				}

				var items []github.Gist

				for {
					page, res, err := app.gh.Gists.ListStarred(opt)
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
			Usage: `get a single gist.`,
			Description: `get a single gist.

   GitHub API docs: http://developer.github.com/v3/gists/#get-a-single-gist`,
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

   GitHub API docs: https://developer.github.com/v3/gists/#get-a-specific-revision-of-a-gist`,
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
			Usage: `create a gist for authenticated user.`,
			Description: `create a gist for authenticated user.

   GitHub API docs: http://developer.github.com/v3/gists/#create-a-gist`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.BoolFlag{Name: `public`, Usage: ``},
				cli.IntFlag{Name: `owner-id`, Usage: ``},
				cli.IntFlag{Name: `owner-public-repos`, Usage: ``},
				cli.StringFlag{Name: `owner-updated-at`, Usage: ``},
				cli.StringFlag{Name: `owner-bio`, Usage: ``},
				cli.BoolFlag{Name: `owner-site-admin`, Usage: ``},
				cli.StringFlag{Name: `owner-location`, Usage: ``},
				cli.IntFlag{Name: `owner-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `owner-hireable`, Usage: ``},
				cli.StringFlag{Name: `owner-login`, Usage: ``},
				cli.StringFlag{Name: `owner-company`, Usage: ``},
				cli.StringFlag{Name: `owner-email`, Usage: ``},
				cli.IntFlag{Name: `owner-following`, Usage: ``},
				cli.StringFlag{Name: `owner-type`, Usage: ``},
				cli.StringFlag{Name: `owner-name`, Usage: ``},
				cli.IntFlag{Name: `owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `owner-followers`, Usage: ``},
				cli.IntFlag{Name: `owner-private-gists`, Usage: ``},
				cli.StringFlag{Name: `owner-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `owner-blog`, Usage: ``},
				cli.StringFlag{Name: `owner-created-at`, Usage: ``},
				cli.IntFlag{Name: `owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-collaborators`, Usage: ``},
				cli.StringFlag{Name: `owner-plan-name`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-space`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `comments`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				gist := &github.Gist{
					Comments:    github.Int(c.Int("comments")),
					CreatedAt:   timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt:   timePointer(now.MustParse(c.String("updated-at"))),
					ID:          github.String(c.String("id")),
					Description: github.String(c.String("description")),
					Public:      github.Bool(c.Bool("public")),
				}

				result, res, err := app.gh.Gists.Create(gist)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit a gist.`,
			Description: `edit a gist.

   GitHub API docs: http://developer.github.com/v3/gists/#edit-a-gist`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.BoolFlag{Name: `public`, Usage: ``},
				cli.BoolFlag{Name: `owner-hireable`, Usage: ``},
				cli.StringFlag{Name: `owner-email`, Usage: ``},
				cli.IntFlag{Name: `owner-following`, Usage: ``},
				cli.StringFlag{Name: `owner-type`, Usage: ``},
				cli.StringFlag{Name: `owner-login`, Usage: ``},
				cli.StringFlag{Name: `owner-company`, Usage: ``},
				cli.IntFlag{Name: `owner-private-gists`, Usage: ``},
				cli.StringFlag{Name: `owner-name`, Usage: ``},
				cli.IntFlag{Name: `owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `owner-followers`, Usage: ``},
				cli.IntFlag{Name: `owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-space`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `owner-plan-name`, Usage: ``},
				cli.StringFlag{Name: `owner-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `owner-blog`, Usage: ``},
				cli.StringFlag{Name: `owner-created-at`, Usage: ``},
				cli.IntFlag{Name: `owner-id`, Usage: ``},
				cli.IntFlag{Name: `owner-public-repos`, Usage: ``},
				cli.StringFlag{Name: `owner-updated-at`, Usage: ``},
				cli.StringFlag{Name: `owner-bio`, Usage: ``},
				cli.BoolFlag{Name: `owner-site-admin`, Usage: ``},
				cli.StringFlag{Name: `owner-location`, Usage: ``},
				cli.IntFlag{Name: `owner-disk-usage`, Usage: ``},
				cli.IntFlag{Name: `comments`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "edit", "edit <id>")
				}

				id := c.Args().Get(0)
				gist := &github.Gist{
					ID:          github.String(c.String("id")),
					Description: github.String(c.String("description")),
					Public:      github.Bool(c.Bool("public")),
					Comments:    github.Int(c.Int("comments")),
					CreatedAt:   timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt:   timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Gists.Edit(id, gist)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete",
			Usage: `delete a gist.`,
			Description: `delete a gist.

   GitHub API docs: http://developer.github.com/v3/gists/#delete-a-gist`,
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

   GitHub API docs: http://developer.github.com/v3/gists/#star-a-gist`,
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

   Github API docs: http://developer.github.com/v3/gists/#unstar-a-gist`,
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

   GitHub API docs: http://developer.github.com/v3/gists/#check-if-a-gist-is-starred`,
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

   GitHub API docs: http://developer.github.com/v3/gists/#fork-a-gist`,
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

   GitHub API docs: http://developer.github.com/v3/gists/comments/#list-comments-on-a-gist`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-comments", "list-comments <gist-id>")
				}

				gistID := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/gists/comments/#get-a-single-comment`,
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
			Usage: `create-comment creates a comment for a gist.`,
			Description: `create-comment creates a comment for a gist.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#create-a-comment`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "create-comment", "create-comment <gist-id>")
				}

				gistID := c.Args().Get(0)
				comment := &github.GistComment{
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
				}

				result, res, err := app.gh.Gists.CreateComment(gistID, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-comment",
			Usage: `edit-comment edits an existing gist comment.`,
			Description: `edit-comment edits an existing gist comment.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#edit-a-comment`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "edit-comment", "edit-comment <gist-id> <comment-id>")
				}

				gistID := c.Args().Get(0)
				commentID, err := strconv.Atoi(c.Args().Get(1))
				check(err)
				comment := &github.GistComment{
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
				}

				result, res, err := app.gh.Gists.EditComment(gistID, commentID, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes a gist comment.`,
			Description: `delete-comment deletes a gist comment.

   GitHub API docs: http://developer.github.com/v3/gists/comments/#delete-a-comment`,
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
