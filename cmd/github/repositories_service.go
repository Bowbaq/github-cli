package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/kr/pretty"
)

var RepositoriesService = cli.Command{
	Name:     "repositories",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "list",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-by-org",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-all",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-contributors",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-languages",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-teams",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-teams <owner> <repo>")
				}

				var items []github.Team

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListTeams(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "list-tags",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-tags <owner> <repo>")
				}

				var items []github.RepositoryTag

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListTags(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "list-branches",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-branches <owner> <repo>")
				}

				var items []github.Branch

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListBranches(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "get-branch",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-collaborators",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-collaborators <owner> <repo>")
				}

				var items []github.User

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListCollaborators(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "is-collaborator",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "add-collaborator",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "remove-collaborator",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-comments",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-comments <owner> <repo>")
				}

				var items []github.RepositoryComment

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListComments(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "list-commit-comments",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-commit-comments <owner> <repo> <sha>")
				}

				var items []github.RepositoryComment

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListCommitComments(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2), opt)
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
			Name: "create-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "update-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-comment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-commits",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-commit",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "compare-commits",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-readme",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "download-contents",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-contents",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-file",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "update-file",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-file",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-archive-link",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-deployments",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-deployment",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-deployment-statuses",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-deployment-statuses <owner> <repo> <deployment>")
				}

				var items []github.DeploymentStatus

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListDeploymentStatuses(c.Args().Get(0), c.Args().Get(1), parseInt(c.Args().Get(2)), opt)
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
			Name: "create-deployment-status",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-forks",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-fork",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-hooks",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-hooks <owner> <repo>")
				}

				var items []github.Hook

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListHooks(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "get-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "ping-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "test-hook",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-service-hooks",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-keys",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-keys <owner> <repo>")
				}

				var items []github.Key

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListKeys(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "get-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-key",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "merge",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-pages-info",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-pages-builds",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-latest-pages-build",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-releases",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + " list-releases <owner> <repo>")
				}

				var items []github.RepositoryRelease

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListReleases(c.Args().Get(0), c.Args().Get(1), opt)
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
			Name: "get-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-latest-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-release-by-tag",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-single-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "create-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-release",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-release-assets",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-release-assets <owner> <repo> <id>")
				}

				var items []github.ReleaseAsset

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListReleaseAssets(c.Args().Get(0), c.Args().Get(1), parseInt(c.Args().Get(2)), opt)
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
			Name: "get-release-asset",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "edit-release-asset",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "delete-release-asset",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "upload-release-asset",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-contributors-stats",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-commit-activity",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-code-frequency",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-participation",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-punch-card",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "list-statuses",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "fetch all the pages",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 0,
					Usage: "fetch this specific page",
				},
				cli.IntFlag{
					Name:  "page-size, ps",
					Value: 30,
					Usage: "fetch <page-size> items per page",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + " list-statuses <owner> <repo> <ref>")
				}

				var items []github.RepoStatus

				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				for {
					page, res, err := app.gh.Repositories.ListStatuses(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2), opt)
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
			Name: "create-status",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name: "get-combined-status",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, RepositoriesService)
}
