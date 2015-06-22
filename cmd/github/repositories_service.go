package main

import (
	"fmt"
	"strconv"

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
			Name:  "list-all",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get",
			Usage: `get fetches a repository.`,
			Description: `get fetches a repository.

   GitHub API docs: http://developer.github.com/v3/repos/#get
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.Get(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete",
			Usage: `delete a repository.`,
			Description: `delete a repository.

   GitHub API docs: https://developer.github.com/v3/repos/#delete-a-repository
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "delete <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				res, err := app.gh.Repositories.Delete(owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-contributors",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-languages",
			Usage: `list-languages lists languages for the specified repository.`,
			Description: `list-languages lists languages for the specified repository. The returned map
   specifies the languages and the number of bytes of code written in that
   language. For example:

       {
         "C": 78769,
         "Python": 7769
       }

   GitHub API Docs: http://developer.github.com/v3/repos/#list-languages
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-languages <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListLanguages(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-teams",
			Usage: `list-teams lists the teams for the specified repository.`,
			Description: `list-teams lists the teams for the specified repository.

   GitHub API docs: https://developer.github.com/v3/repos/#list-teams
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-teams <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Team

				for {
					page, res, err := app.gh.Repositories.ListTeams(owner, repo, opt)
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
			Name:  "list-tags",
			Usage: `list-tags lists tags for the specified repository.`,
			Description: `list-tags lists tags for the specified repository.

   GitHub API docs: https://developer.github.com/v3/repos/#list-tags
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-tags <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepositoryTag

				for {
					page, res, err := app.gh.Repositories.ListTags(owner, repo, opt)
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
			Name:  "list-branches",
			Usage: `list-branches lists branches for the specified repository.`,
			Description: `list-branches lists branches for the specified repository.

   GitHub API docs: http://developer.github.com/v3/repos/#list-branches
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-branches <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Branch

				for {
					page, res, err := app.gh.Repositories.ListBranches(owner, repo, opt)
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
			Name:  "get-branch",
			Usage: `get-branch gets the specified branch for a repository.`,
			Description: `get-branch gets the specified branch for a repository.

   GitHub API docs: https://developer.github.com/v3/repos/#get-branch
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-branch <owner> <repo> <branch>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				branch := c.Args().Get(2)

				result, res, err := app.gh.Repositories.GetBranch(owner, repo, branch)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-collaborators",
			Usage: `list-collaborators lists the Github users that have access to the repository.`,
			Description: `list-collaborators lists the Github users that have access to the repository.

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#list
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-collaborators <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Repositories.ListCollaborators(owner, repo, opt)
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
			Name:  "is-collaborator",
			Usage: `is-collaborator checks whether the specified Github user has collaborator access to the given repo.`,
			Description: `is-collaborator checks whether the specified Github user has collaborator
   access to the given repo.
   Note: This will return false if the user is not a collaborator OR the user
   is not a GitHub user.

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#get
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "is-collaborator <owner> <repo> <user>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				user := c.Args().Get(2)

				result, res, err := app.gh.Repositories.IsCollaborator(owner, repo, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "add-collaborator",
			Usage: `add-collaborator adds the specified Github user as collaborator to the given repo.`,
			Description: `add-collaborator adds the specified Github user as collaborator to the given repo.

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#add-collaborator
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "add-collaborator <owner> <repo> <user>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				user := c.Args().Get(2)

				res, err := app.gh.Repositories.AddCollaborator(owner, repo, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "remove-collaborator",
			Usage: `remove-collaborator removes the specified Github user as collaborator from the given repo.`,
			Description: `remove-collaborator removes the specified Github user as collaborator from the given repo.
   Note: Does not return error if a valid user that is not a collaborator is removed.

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#remove-collaborator
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "remove-collaborator <owner> <repo> <user>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				user := c.Args().Get(2)

				res, err := app.gh.Repositories.RemoveCollaborator(owner, repo, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-comments",
			Usage: `list-comments lists all the comments for the repository.`,
			Description: `list-comments lists all the comments for the repository.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-comments <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepositoryComment

				for {
					page, res, err := app.gh.Repositories.ListComments(owner, repo, opt)
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
			Name:  "list-commit-comments",
			Usage: `list-commit-comments lists all the comments for a given commit SHA.`,
			Description: `list-commit-comments lists all the comments for a given commit SHA.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-commit-comments <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepositoryComment

				for {
					page, res, err := app.gh.Repositories.ListCommitComments(owner, repo, sha, opt)
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
			Name:  "create-comment",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-comment",
			Usage: `get-comment gets a single comment from a repository.`,
			Description: `get-comment gets a single comment from a repository.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#get-a-single-commit-comment
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

				result, res, err := app.gh.Repositories.GetComment(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "update-comment",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes a single comment from a repository.`,
			Description: `delete-comment deletes a single comment from a repository.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#delete-a-commit-comment
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

				res, err := app.gh.Repositories.DeleteComment(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-commits",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-commit",
			Usage: `get-commit fetches the specified commit, including all details about it.`,
			Description: `get-commit fetches the specified commit, including all details about it.
   todo: support media formats - https://github.com/google/go-github/issues/6

   GitHub API docs: http://developer.github.com/v3/repos/commits/#get-a-single-commit
   See also: http://developer.github.com//v3/git/commits/#get-a-single-commit provides the same functionality
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-commit <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)

				result, res, err := app.gh.Repositories.GetCommit(owner, repo, sha)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "compare-commits",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-readme",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "download-contents",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-contents",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-file",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "update-file",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-file",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-archive-link",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-deployments",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-deployment",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-deployment-statuses",
			Usage: `list-deployment-statuses lists the statuses of a given deployment of a repository.`,
			Description: `list-deployment-statuses lists the statuses of a given deployment of a repository.

   GitHub API docs: https://developer.github.com/v3/repos/deployments/#list-deployment-statuses
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-deployment-statuses <owner> <repo> <deployment>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				deployment, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.DeploymentStatus

				for {
					page, res, err := app.gh.Repositories.ListDeploymentStatuses(owner, repo, deployment, opt)
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
			Name:  "create-deployment-status",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-forks",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-fork",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-hook",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-hooks",
			Usage: `list-hooks lists all Hooks for the specified repository.`,
			Description: `list-hooks lists all Hooks for the specified repository.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#list
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-hooks <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Hook

				for {
					page, res, err := app.gh.Repositories.ListHooks(owner, repo, opt)
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
			Name:  "get-hook",
			Usage: `get-hook returns a single specified Hook.`,
			Description: `get-hook returns a single specified Hook.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#get-single-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Repositories.GetHook(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-hook",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-hook",
			Usage: `delete-hook deletes a specified Hook.`,
			Description: `delete-hook deletes a specified Hook.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#delete-a-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.DeleteHook(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "ping-hook",
			Usage: `ping-hook triggers a 'ping' event to be sent to the Hook.`,
			Description: `ping-hook triggers a 'ping' event to be sent to the Hook.

   GitHub API docs: https://developer.github.com/v3/repos/hooks/#ping-a-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "ping-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.PingHook(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "test-hook",
			Usage: `test-hook triggers a test Hook by github.`,
			Description: `test-hook triggers a test Hook by github.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#test-a-push-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "test-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.TestHook(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-service-hooks",
			Usage: `list-service-hooks is deprecated.`,
			Description: `list-service-hooks is deprecated.  Use Client.list-service-hooks instead.
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {

				result, res, err := app.gh.Repositories.ListServiceHooks()
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-keys",
			Usage: `list-keys lists the deploy keys for a repository.`,
			Description: `list-keys lists the deploy keys for a repository.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#list
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-keys <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.Key

				for {
					page, res, err := app.gh.Repositories.ListKeys(owner, repo, opt)
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
			Name:  "get-key",
			Usage: `get-key fetches a single deploy key.`,
			Description: `get-key fetches a single deploy key.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#get
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-key <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Repositories.GetKey(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-key",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-key",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-key",
			Usage: `delete-key deletes a deploy key.`,
			Description: `delete-key deletes a deploy key.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#delete
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-key <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.DeleteKey(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "merge",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-pages-info",
			Usage: `get-pages-info fetches information about a GitHub Pages site.`,
			Description: `get-pages-info fetches information about a GitHub Pages site.

   GitHub API docs: https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-pages-info <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.GetPagesInfo(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-pages-builds",
			Usage: `list-pages-builds lists the builds for a GitHub Pages site.`,
			Description: `list-pages-builds lists the builds for a GitHub Pages site.

   GitHub API docs: https://developer.github.com/v3/repos/pages/#list-pages-builds
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-pages-builds <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListPagesBuilds(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-latest-pages-build",
			Usage: `get-latest-pages-build fetches the latest build information for a GitHub pages site.`,
			Description: `get-latest-pages-build fetches the latest build information for a GitHub pages site.

   GitHub API docs: https://developer.github.com/v3/repos/pages/#list-latest-pages-build
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-latest-pages-build <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.GetLatestPagesBuild(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-releases",
			Usage: `list-releases lists the releases for a repository.`,
			Description: `list-releases lists the releases for a repository.

   GitHub API docs: http://developer.github.com/v3/repos/releases/#list-releases-for-a-repository
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-releases <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepositoryRelease

				for {
					page, res, err := app.gh.Repositories.ListReleases(owner, repo, opt)
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
			Name:  "get-release",
			Usage: `get-release fetches a single release.`,
			Description: `get-release fetches a single release.

   GitHub API docs: http://developer.github.com/v3/repos/releases/#get-a-single-release
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-release <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Repositories.GetRelease(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-latest-release",
			Usage: `get-latest-release fetches the latest published release for the repository.`,
			Description: `get-latest-release fetches the latest published release for the repository.

   GitHub API docs: https://developer.github.com/v3/repos/releases/#get-the-latest-release
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-latest-release <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.GetLatestRelease(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-release-by-tag",
			Usage: `GetLatestReleaseByTag fetches a release with the specified tag.`,
			Description: `GetLatestReleaseByTag fetches a release with the specified tag.

   GitHub API docs: https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-release-by-tag <owner> <repo> <tag>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				tag := c.Args().Get(2)

				result, res, err := app.gh.Repositories.GetReleaseByTag(owner, repo, tag)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-release",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-release",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-release",
			Usage: `delete-release delete a single release from a repository.`,
			Description: `delete-release delete a single release from a repository.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#delete-a-release
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-release <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.DeleteRelease(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-release-assets",
			Usage: `list-release-assets lists the release's assets.`,
			Description: `list-release-assets lists the release's assets.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#list-assets-for-a-release
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-release-assets <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.ReleaseAsset

				for {
					page, res, err := app.gh.Repositories.ListReleaseAssets(owner, repo, id, opt)
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
			Name:  "get-release-asset",
			Usage: `get-release-asset fetches a single release asset.`,
			Description: `get-release-asset fetches a single release asset.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#get-a-single-release-asset
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "get-release-asset <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				result, res, err := app.gh.Repositories.GetReleaseAsset(owner, repo, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-release-asset",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-release-asset",
			Usage: `delete-release-asset delete a single release asset from a repository.`,
			Description: `delete-release-asset delete a single release asset from a repository.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#delete-a-release-asset
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "delete-release-asset <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.DeleteReleaseAsset(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "upload-release-asset",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-contributors-stats",
			Usage: `list-contributors-stats gets a repo's contributor list with additions, deletions and commit counts.`,
			Description: `list-contributors-stats gets a repo's contributor list with additions,
   deletions and commit counts.

   If this is the first time these statistics are requested for the given
   repository, this method will return a non-nil error and a status code of
   202. This is because this is the status that github returns to signify that
   it is now computing the requested statistics. A follow up request, after a
   delay of a second or so, should result in a successful request.

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#contributors
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-contributors-stats <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListContributorsStats(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-commit-activity",
			Usage: `list-commit-activity returns the last year of commit activity grouped by week.`,
			Description: `list-commit-activity returns the last year of commit activity
   grouped by week. The days array is a group of commits per day,
   starting on Sunday.

   If this is the first time these statistics are requested for the given
   repository, this method will return a non-nil error and a status code of
   202. This is because this is the status that github returns to signify that
   it is now computing the requested statistics. A follow up request, after a
   delay of a second or so, should result in a successful request.

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#commit-activity
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-commit-activity <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListCommitActivity(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-code-frequency",
			Usage: `list-code-frequency returns a weekly aggregate of the number of additions and deletions pushed to a repository.`,
			Description: `list-code-frequency returns a weekly aggregate of the number of additions and
   deletions pushed to a repository.  Returned WeeklyStats will contain
   additiona and deletions, but not total commits.

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#code-frequency
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-code-frequency <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListCodeFrequency(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-participation",
			Usage: `list-participation returns the total commit counts for the 'owner' and total commit counts in 'all'.`,
			Description: `list-participation returns the total commit counts for the 'owner'
   and total commit counts in 'all'. 'all' is everyone combined,
   including the 'owner' in the last 52 weeks. If you’d like to get
   the commit counts for non-owners, you can subtract 'all' from 'owner'.

   The array order is oldest week (index 0) to most recent week.

   If this is the first time these statistics are requested for the given
   repository, this method will return a non-nil error and a status code
   of 202. This is because this is the status that github returns to
   signify that it is now computing the requested statistics. A follow
   up request, after a delay of a second or so, should result in a
   successful request.

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#participation
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-participation <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListParticipation(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-punch-card",
			Usage: `list-punch-card returns the number of commits per hour in each day.`,
			Description: `list-punch-card returns the number of commits per hour in each day.

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#punch-card
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "list-punch-card <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.ListPunchCard(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-statuses",
			Usage: `list-statuses lists the statuses of a repository at the specified reference.`,
			Description: `list-statuses lists the statuses of a repository at the specified
   reference.  ref can be a SHA, a branch name, or a tag name.

   GitHub API docs: http://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "list-statuses <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
				}

				var items []github.RepoStatus

				for {
					page, res, err := app.gh.Repositories.ListStatuses(owner, repo, ref, opt)
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
			Name:  "create-status",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-combined-status",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, RepositoriesService)
}
