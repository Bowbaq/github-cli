package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var SearchService = cli.Command{
	Name:     "search",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "repositories",
			Usage: `repositories searches repositories via various criteria.`,
			Description: `repositories searches repositories via various criteria.

   GitHub API docs: http://developer.github.com/v3/search/#search-repositories`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `text-match`, Usage: `Whether to retrieve text match metadata with a query`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `sort`, Usage: `How to sort the search results.  Possible values are:
  - for repositories: stars, fork, updated
  - for code: indexed
  - for issues: comments, created, updated
  - for users: followers, repositories, joined

Default is to sort by best match.`},
				cli.StringFlag{Name: `order`, Usage: `Sort order if sort parameter is provided. Possible values are: asc,
desc. Default is desc.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "repositories", "repositories <query>")
				}

				query := c.Args().Get(0)
				opt := &github.SearchOptions{
					Sort:      c.String("sort"),
					Order:     c.String("order"),
					TextMatch: c.Bool("text-match"),
				}

				result, res, err := app.gh.Search.Repositories(query, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "issues",
			Usage: `issues searches issues via various criteria.`,
			Description: `issues searches issues via various criteria.

   GitHub API docs: http://developer.github.com/v3/search/#search-issues`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sort`, Usage: `How to sort the search results.  Possible values are:
  - for repositories: stars, fork, updated
  - for code: indexed
  - for issues: comments, created, updated
  - for users: followers, repositories, joined

Default is to sort by best match.`},
				cli.StringFlag{Name: `order`, Usage: `Sort order if sort parameter is provided. Possible values are: asc,
desc. Default is desc.`},
				cli.BoolFlag{Name: `text-match`, Usage: `Whether to retrieve text match metadata with a query`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "issues", "issues <query>")
				}

				query := c.Args().Get(0)
				opt := &github.SearchOptions{
					Sort:      c.String("sort"),
					Order:     c.String("order"),
					TextMatch: c.Bool("text-match"),
				}

				result, res, err := app.gh.Search.Issues(query, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "users",
			Usage: `users searches users via various criteria.`,
			Description: `users searches users via various criteria.

   GitHub API docs: http://developer.github.com/v3/search/#search-users`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `order`, Usage: `Sort order if sort parameter is provided. Possible values are: asc,
desc. Default is desc.`},
				cli.BoolFlag{Name: `text-match`, Usage: `Whether to retrieve text match metadata with a query`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `sort`, Usage: `How to sort the search results.  Possible values are:
  - for repositories: stars, fork, updated
  - for code: indexed
  - for issues: comments, created, updated
  - for users: followers, repositories, joined

Default is to sort by best match.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "users", "users <query>")
				}

				query := c.Args().Get(0)
				opt := &github.SearchOptions{
					Sort:      c.String("sort"),
					Order:     c.String("order"),
					TextMatch: c.Bool("text-match"),
				}

				result, res, err := app.gh.Search.Users(query, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "code",
			Usage: `code searches code via various criteria.`,
			Description: `code searches code via various criteria.

   GitHub API docs: http://developer.github.com/v3/search/#search-code`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `text-match`, Usage: `Whether to retrieve text match metadata with a query`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `sort`, Usage: `How to sort the search results.  Possible values are:
  - for repositories: stars, fork, updated
  - for code: indexed
  - for issues: comments, created, updated
  - for users: followers, repositories, joined

Default is to sort by best match.`},
				cli.StringFlag{Name: `order`, Usage: `Sort order if sort parameter is provided. Possible values are: asc,
desc. Default is desc.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "code", "code <query>")
				}

				query := c.Args().Get(0)
				opt := &github.SearchOptions{
					Order:     c.String("order"),
					TextMatch: c.Bool("text-match"),
					Sort:      c.String("sort"),
				}

				result, res, err := app.gh.Search.Code(query, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, SearchService)
}
