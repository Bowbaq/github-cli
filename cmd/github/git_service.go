package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/kr/pretty"
)

var GitService = cli.Command{
	Name:     "git",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "get-blob",
			Usage: `get-blob fetchs a blob from a repo given a SHA.`,
			Description: `get-blob fetchs a blob from a repo given a SHA.

   GitHub API docs: http://developer.github.com/v3/git/blobs/#get-a-blob
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-blob", "get-blob <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)

				result, res, err := app.gh.Git.GetBlob(owner, repo, sha)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-blob",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-commit",
			Usage: `get-commit fetchs the Commit object for a given SHA.`,
			Description: `get-commit fetchs the Commit object for a given SHA.

   GitHub API docs: http://developer.github.com/v3/git/commits/#get-a-commit
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-commit", "get-commit <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)

				result, res, err := app.gh.Git.GetCommit(owner, repo, sha)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-commit",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-ref",
			Usage: `get-ref fetches the Reference object for a given Git ref.`,
			Description: `get-ref fetches the Reference object for a given Git ref.

   GitHub API docs: http://developer.github.com/v3/git/refs/#get-a-reference
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-ref", "get-ref <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)

				result, res, err := app.gh.Git.GetRef(owner, repo, ref)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-refs",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-ref",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "update-ref",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-ref",
			Usage: `delete-ref deletes a ref from a repository.`,
			Description: `delete-ref deletes a ref from a repository.

   GitHub API docs: http://developer.github.com/v3/git/refs/#delete-a-reference
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-ref", "delete-ref <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)

				res, err := app.gh.Git.DeleteRef(owner, repo, ref)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "get-tag",
			Usage: `get-tag fetchs a tag from a repo given a SHA.`,
			Description: `get-tag fetchs a tag from a repo given a SHA.

   GitHub API docs: http://developer.github.com/v3/git/tags/#get-a-tag
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-tag", "get-tag <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)

				result, res, err := app.gh.Git.GetTag(owner, repo, sha)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-tag",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-tree",
			Usage: `get-tree fetches the Tree object for a given sha hash from a repository.`,
			Description: `get-tree fetches the Tree object for a given sha hash from a repository.

   GitHub API docs: http://developer.github.com/v3/git/trees/#get-a-tree
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "recursive"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					showHelp(c, "get-tree", "get-tree <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)
				recursive := c.Bool("recursive")

				result, res, err := app.gh.Git.GetTree(owner, repo, sha, recursive)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-tree",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, GitService)
}
