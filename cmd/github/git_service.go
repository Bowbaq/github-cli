package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
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

   GitHub API docs: http://developer.github.com/v3/git/blobs/#get-a-blob`,
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
			Usage: `create-blob creates a blob object.`,
			Description: `create-blob creates a blob object.

   GitHub API docs: http://developer.github.com/v3/git/blobs/#create-a-blob`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `encoding`, Usage: ``},
				cli.StringFlag{Name: `sha`, Usage: ``},
				cli.IntFlag{Name: `size`, Usage: ``},
				cli.StringFlag{Name: `content`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-blob", "create-blob <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				blob := &github.Blob{
					Content:  github.String(c.String("content")),
					Encoding: github.String(c.String("encoding")),
					SHA:      github.String(c.String("sha")),
					Size:     github.Int(c.Int("size")),
				}

				result, res, err := app.gh.Git.CreateBlob(owner, repo, blob)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-commit",
			Usage: `get-commit fetchs the Commit object for a given SHA.`,
			Description: `get-commit fetchs the Commit object for a given SHA.

   GitHub API docs: http://developer.github.com/v3/git/commits/#get-a-commit`,
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
			Usage: `create-commit creates a new commit in a repository.`,
			Description: `create-commit creates a new commit in a repository.

   The commit.Committer is optional and will be filled with the commit.Author
   data if omitted. If the commit.Author is omitted, it will be filled in with
   the authenticated user’s information and the current date.

   GitHub API docs: http://developer.github.com/v3/git/commits/#create-a-commit`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `author-email`, Usage: ``},
				cli.StringFlag{Name: `author-date`, Usage: ``},
				cli.StringFlag{Name: `author-name`, Usage: ``},
				cli.StringFlag{Name: `tree-sha`, Usage: ``},
				cli.StringFlag{Name: `message`, Usage: ``},
				cli.IntFlag{Name: `stats-deletions`, Usage: ``},
				cli.IntFlag{Name: `stats-total`, Usage: ``},
				cli.IntFlag{Name: `stats-additions`, Usage: ``},
				cli.IntFlag{Name: `comment-count`, Usage: `CommentCount is the number of GitHub comments on the commit.  This
is only populated for requests that fetch GitHub data like
Pulls.ListCommits, Repositories.ListCommits, etc.`},
				cli.StringFlag{Name: `sha`, Usage: ``},
				cli.StringFlag{Name: `committer-email`, Usage: ``},
				cli.StringFlag{Name: `committer-date`, Usage: ``},
				cli.StringFlag{Name: `committer-name`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-commit", "create-commit <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				commit := &github.Commit{
					CommentCount: github.Int(c.Int("comment-count")),
					SHA:          github.String(c.String("sha")),
					Message:      github.String(c.String("message")),
				}

				result, res, err := app.gh.Git.CreateCommit(owner, repo, commit)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-ref",
			Usage: `get-ref fetches the Reference object for a given Git ref.`,
			Description: `get-ref fetches the Reference object for a given Git ref.

   GitHub API docs: http://developer.github.com/v3/git/refs/#get-a-reference`,
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
			Usage: `list-refs lists all refs in a repository.`,
			Description: `list-refs lists all refs in a repository.

   GitHub API docs: http://developer.github.com/v3/git/refs/#get-all-references`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `type`, Usage: ``},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-refs", "list-refs <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ReferenceListOptions{
					Type: c.String("type"),
				}

				var items []github.Reference

				for {
					page, res, err := app.gh.Git.ListRefs(owner, repo, opt)
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
			Name:  "create-ref",
			Usage: `create-ref creates a new ref in a repository.`,
			Description: `create-ref creates a new ref in a repository.

   GitHub API docs: http://developer.github.com/v3/git/refs/#create-a-reference`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `ref`, Usage: ``},
				cli.StringFlag{Name: `object-sha`, Usage: ``},
				cli.StringFlag{Name: `object-type`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-ref", "create-ref <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := &github.Reference{
					Ref: github.String(c.String("ref")),
				}

				result, res, err := app.gh.Git.CreateRef(owner, repo, ref)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "update-ref",
			Usage: `update-ref updates an existing ref in a repository.`,
			Description: `update-ref updates an existing ref in a repository.

   GitHub API docs: http://developer.github.com/v3/git/refs/#update-a-reference`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `ref`, Usage: ``},
				cli.StringFlag{Name: `object-type`, Usage: ``},
				cli.StringFlag{Name: `object-sha`, Usage: ``},
				cli.BoolFlag{Name: `force`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "update-ref", "update-ref <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := &github.Reference{
					Ref: github.String(c.String("ref")),
				}
				force := c.Bool("force")

				result, res, err := app.gh.Git.UpdateRef(owner, repo, ref, force)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-ref",
			Usage: `delete-ref deletes a ref from a repository.`,
			Description: `delete-ref deletes a ref from a repository.

   GitHub API docs: http://developer.github.com/v3/git/refs/#delete-a-reference`,
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

   GitHub API docs: http://developer.github.com/v3/git/tags/#get-a-tag`,
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
			Usage: `create-tag creates a tag object.`,
			Description: `create-tag creates a tag object.

   GitHub API docs: http://developer.github.com/v3/git/tags/#create-a-tag-object`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `message`, Usage: ``},
				cli.StringFlag{Name: `tagger-date`, Usage: ``},
				cli.StringFlag{Name: `tagger-name`, Usage: ``},
				cli.StringFlag{Name: `tagger-email`, Usage: ``},
				cli.StringFlag{Name: `object-type`, Usage: ``},
				cli.StringFlag{Name: `object-sha`, Usage: ``},
				cli.StringFlag{Name: `tag`, Usage: ``},
				cli.StringFlag{Name: `sha`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-tag", "create-tag <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				tag := &github.Tag{
					Tag:     github.String(c.String("tag")),
					SHA:     github.String(c.String("sha")),
					Message: github.String(c.String("message")),
				}

				result, res, err := app.gh.Git.CreateTag(owner, repo, tag)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-tree",
			Usage: `get-tree fetches the Tree object for a given sha hash from a repository.`,
			Description: `get-tree fetches the Tree object for a given sha hash from a repository.

   GitHub API docs: http://developer.github.com/v3/git/trees/#get-a-tree`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `recursive`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
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
			Usage: `create-tree creates a new tree in a repository.`,
			Description: `create-tree creates a new tree in a repository.  If both a tree and a nested
   path modifying that tree are specified, it will overwrite the contents of
   that tree with the new path contents and write a new tree out.

   GitHub API docs: http://developer.github.com/v3/git/trees/#create-a-tree`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-tree", "create-tree <owner> <repo> <base-tree>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				baseTree := c.Args().Get(2)
				var entries []github.TreeEntry

				result, res, err := app.gh.Git.CreateTree(owner, repo, baseTree, entries)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, GitService)
}
