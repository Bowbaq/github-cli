package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
	"github.com/kr/pretty"
)

var RepositoriesService = cli.Command{
	Name:     "repositories",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list the repositories for a user.`,
			Description: `list the repositories for a user.  Passing the empty string will list
   repositories for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/repos/#list-user-repositories`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `type`, Usage: `Type of repositories to list.  Possible values are: all, owner, public,
private, member.  Default is "all".`},
				cli.StringFlag{Name: `sort`, Usage: `How to sort the repository list.  Possible values are: created, updated,
pushed, full_name.  Default is "full_name".`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort repositories.  Possible values are: asc, desc.
Default is "asc" when sort is "full_name", otherwise default is "desc".`},
				cli.BoolFlag{Name: `include-org`, Usage: `Include orginization repositories the user has access to.
This will become the default behavior in the future, but is opt-in for now.
See https://developer.github.com/changes/2015-01-07-prepare-for-organization-permissions-changes/`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list", "list <user>")
				}

				user := c.Args().Get(0)
				opt := &github.RepositoryListOptions{
					Type:       c.String("type"),
					Sort:       c.String("sort"),
					Direction:  c.String("direction"),
					IncludeOrg: c.Bool("include-org"),
				}

				var items []github.Repository

				for {
					page, res, err := app.gh.Repositories.List(user, opt)
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
			Usage: `list-by-org lists the repositories for an organization.`,
			Description: `list-by-org lists the repositories for an organization.

   GitHub API docs: http://developer.github.com/v3/repos/#list-organization-repositories`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `type`, Usage: `Type of repositories to list.  Possible values are: all, public, private,
forks, sources, member.  Default is "all".`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-by-org", "list-by-org <org>")
				}

				org := c.Args().Get(0)
				opt := &github.RepositoryListByOrgOptions{
					Type: c.String("type"),
				}

				var items []github.Repository

				for {
					page, res, err := app.gh.Repositories.ListByOrg(org, opt)
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
			Usage: `list-all lists all GitHub repositories in the order that they were created.`,
			Description: `list-all lists all GitHub repositories in the order that they were created.

   GitHub API docs: http://developer.github.com/v3/repos/#list-all-public-repositories`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `since`, Usage: `ID of the last repository seen`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.RepositoryListAllOptions{
					Since: c.Int("since"),
				}

				var items []github.Repository

				for {
					page, res, err := app.gh.Repositories.ListAll(opt)
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
			Name:  "create",
			Usage: `create a new repository.`,
			Description: `create a new repository.  If an organization is specified, the new
   repository will be created under that org.  If the empty string is
   specified, it will be created for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/repos/#create`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `full-name`, Usage: ``},
				cli.BoolFlag{Name: `has-downloads`, Usage: ``},
				cli.IntFlag{Name: `network-count`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.IntFlag{Name: `open-issues-count`, Usage: ``},
				cli.StringFlag{Name: `organization-created-at`, Usage: ``},
				cli.StringFlag{Name: `organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-id`, Usage: ``},
				cli.StringFlag{Name: `organization-location`, Usage: ``},
				cli.StringFlag{Name: `organization-email`, Usage: ``},
				cli.IntFlag{Name: `organization-public-gists`, Usage: ``},
				cli.StringFlag{Name: `organization-updated-at`, Usage: ``},
				cli.IntFlag{Name: `organization-total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-type`, Usage: ``},
				cli.StringFlag{Name: `organization-company`, Usage: ``},
				cli.StringFlag{Name: `organization-blog`, Usage: ``},
				cli.IntFlag{Name: `organization-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-public-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-following`, Usage: ``},
				cli.StringFlag{Name: `organization-name`, Usage: ``},
				cli.IntFlag{Name: `organization-disk-usage`, Usage: ``},
				cli.IntFlag{Name: `organization-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-followers`, Usage: ``},
				cli.IntFlag{Name: `organization-private-gists`, Usage: ``},
				cli.StringFlag{Name: `organization-login`, Usage: ``},
				cli.StringFlag{Name: `organization-billing-email`, Usage: ``},
				cli.BoolFlag{Name: `private`, Usage: `Additional mutable fields when creating and editing a repository`},
				cli.BoolFlag{Name: `has-issues`, Usage: ``},
				cli.IntFlag{Name: `owner-private-gists`, Usage: ``},
				cli.StringFlag{Name: `owner-name`, Usage: ``},
				cli.IntFlag{Name: `owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `owner-followers`, Usage: ``},
				cli.IntFlag{Name: `owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-collaborators`, Usage: ``},
				cli.StringFlag{Name: `owner-plan-name`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-space`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-private-repos`, Usage: ``},
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
				cli.BoolFlag{Name: `owner-hireable`, Usage: ``},
				cli.StringFlag{Name: `owner-email`, Usage: ``},
				cli.IntFlag{Name: `owner-following`, Usage: ``},
				cli.StringFlag{Name: `owner-type`, Usage: ``},
				cli.StringFlag{Name: `owner-login`, Usage: ``},
				cli.StringFlag{Name: `owner-company`, Usage: ``},
				cli.StringFlag{Name: `homepage`, Usage: ``},
				cli.StringFlag{Name: `license-implementation`, Usage: ``},
				cli.StringSliceFlag{Name: `license-forbidden`, Usage: ``},
				cli.StringFlag{Name: `license-key`, Usage: ``},
				cli.StringFlag{Name: `license-category`, Usage: ``},
				cli.StringSliceFlag{Name: `license-permitted`, Usage: ``},
				cli.StringFlag{Name: `license-body`, Usage: ``},
				cli.StringFlag{Name: `license-name`, Usage: ``},
				cli.BoolFlag{Name: `license-featured`, Usage: ``},
				cli.StringFlag{Name: `license-description`, Usage: ``},
				cli.StringSliceFlag{Name: `license-required`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.BoolFlag{Name: `auto-init`, Usage: ``},
				cli.IntFlag{Name: `team-id`, Usage: `Creating an organization repository. Required for non-owners.`},
				cli.StringFlag{Name: `master-branch`, Usage: ``},
				cli.IntFlag{Name: `stargazers-count`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.BoolFlag{Name: `has-wiki`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.IntFlag{Name: `size`, Usage: ``},
				cli.StringFlag{Name: `default-branch`, Usage: ``},
				cli.IntFlag{Name: `forks-count`, Usage: ``},
				cli.StringFlag{Name: `pushed-at`, Usage: ``},
				cli.BoolFlag{Name: `fork`, Usage: ``},
				cli.IntFlag{Name: `subscribers-count`, Usage: ``},
				cli.StringFlag{Name: `language`, Usage: ``},
				cli.IntFlag{Name: `watchers-count`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "create", "create <org>")
				}

				org := c.Args().Get(0)
				repo := &github.Repository{
					FullName:         github.String(c.String("full-name")),
					NetworkCount:     github.Int(c.Int("network-count")),
					HasDownloads:     github.Bool(c.Bool("has-downloads")),
					ID:               github.Int(c.Int("id")),
					Private:          github.Bool(c.Bool("private")),
					HasIssues:        github.Bool(c.Bool("has-issues")),
					OpenIssuesCount:  github.Int(c.Int("open-issues-count")),
					Homepage:         github.String(c.String("homepage")),
					Description:      github.String(c.String("description")),
					TeamID:           github.Int(c.Int("team-id")),
					MasterBranch:     github.String(c.String("master-branch")),
					AutoInit:         github.Bool(c.Bool("auto-init")),
					StargazersCount:  github.Int(c.Int("stargazers-count")),
					UpdatedAt:        &github.Timestamp{now.MustParse(c.String("updated-at"))},
					CreatedAt:        &github.Timestamp{now.MustParse(c.String("created-at"))},
					HasWiki:          github.Bool(c.Bool("has-wiki")),
					DefaultBranch:    github.String(c.String("default-branch")),
					Size:             github.Int(c.Int("size")),
					PushedAt:         &github.Timestamp{now.MustParse(c.String("pushed-at"))},
					ForksCount:       github.Int(c.Int("forks-count")),
					Language:         github.String(c.String("language")),
					Fork:             github.Bool(c.Bool("fork")),
					SubscribersCount: github.Int(c.Int("subscribers-count")),
					WatchersCount:    github.Int(c.Int("watchers-count")),
					Name:             github.String(c.String("name")),
				}

				result, res, err := app.gh.Repositories.Create(org, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get",
			Usage: `get fetches a repository.`,
			Description: `get fetches a repository.

   GitHub API docs: http://developer.github.com/v3/repos/#get`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get", "get <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				result, res, err := app.gh.Repositories.Get(owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit updates a repository.`,
			Description: `edit updates a repository.

   GitHub API docs: http://developer.github.com/v3/repos/#edit`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `master-branch`, Usage: ``},
				cli.BoolFlag{Name: `auto-init`, Usage: ``},
				cli.IntFlag{Name: `team-id`, Usage: `Creating an organization repository. Required for non-owners.`},
				cli.IntFlag{Name: `stargazers-count`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.BoolFlag{Name: `has-wiki`, Usage: ``},
				cli.StringFlag{Name: `default-branch`, Usage: ``},
				cli.IntFlag{Name: `size`, Usage: ``},
				cli.StringFlag{Name: `pushed-at`, Usage: ``},
				cli.IntFlag{Name: `forks-count`, Usage: ``},
				cli.StringFlag{Name: `language`, Usage: ``},
				cli.BoolFlag{Name: `fork`, Usage: ``},
				cli.IntFlag{Name: `subscribers-count`, Usage: ``},
				cli.IntFlag{Name: `watchers-count`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `full-name`, Usage: ``},
				cli.IntFlag{Name: `network-count`, Usage: ``},
				cli.BoolFlag{Name: `has-downloads`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.BoolFlag{Name: `has-issues`, Usage: ``},
				cli.StringFlag{Name: `owner-bio`, Usage: ``},
				cli.BoolFlag{Name: `owner-site-admin`, Usage: ``},
				cli.StringFlag{Name: `owner-location`, Usage: ``},
				cli.IntFlag{Name: `owner-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `owner-hireable`, Usage: ``},
				cli.IntFlag{Name: `owner-following`, Usage: ``},
				cli.StringFlag{Name: `owner-type`, Usage: ``},
				cli.StringFlag{Name: `owner-login`, Usage: ``},
				cli.StringFlag{Name: `owner-company`, Usage: ``},
				cli.StringFlag{Name: `owner-email`, Usage: ``},
				cli.StringFlag{Name: `owner-name`, Usage: ``},
				cli.IntFlag{Name: `owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `owner-followers`, Usage: ``},
				cli.IntFlag{Name: `owner-private-gists`, Usage: ``},
				cli.IntFlag{Name: `owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `owner-plan-name`, Usage: ``},
				cli.IntFlag{Name: `owner-plan-space`, Usage: ``},
				cli.StringFlag{Name: `owner-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `owner-blog`, Usage: ``},
				cli.StringFlag{Name: `owner-created-at`, Usage: ``},
				cli.IntFlag{Name: `owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `owner-id`, Usage: ``},
				cli.IntFlag{Name: `owner-public-repos`, Usage: ``},
				cli.StringFlag{Name: `owner-updated-at`, Usage: ``},
				cli.IntFlag{Name: `open-issues-count`, Usage: ``},
				cli.StringFlag{Name: `organization-created-at`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `organization-id`, Usage: ``},
				cli.StringFlag{Name: `organization-location`, Usage: ``},
				cli.IntFlag{Name: `organization-public-gists`, Usage: ``},
				cli.StringFlag{Name: `organization-updated-at`, Usage: ``},
				cli.IntFlag{Name: `organization-total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-type`, Usage: ``},
				cli.StringFlag{Name: `organization-company`, Usage: ``},
				cli.StringFlag{Name: `organization-blog`, Usage: ``},
				cli.StringFlag{Name: `organization-email`, Usage: ``},
				cli.IntFlag{Name: `organization-public-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-following`, Usage: ``},
				cli.IntFlag{Name: `organization-owned-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-name`, Usage: ``},
				cli.IntFlag{Name: `organization-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-followers`, Usage: ``},
				cli.IntFlag{Name: `organization-private-gists`, Usage: ``},
				cli.IntFlag{Name: `organization-disk-usage`, Usage: ``},
				cli.StringFlag{Name: `organization-login`, Usage: ``},
				cli.StringFlag{Name: `organization-billing-email`, Usage: ``},
				cli.BoolFlag{Name: `private`, Usage: `Additional mutable fields when creating and editing a repository`},
				cli.StringFlag{Name: `homepage`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.StringFlag{Name: `license-body`, Usage: ``},
				cli.StringFlag{Name: `license-name`, Usage: ``},
				cli.BoolFlag{Name: `license-featured`, Usage: ``},
				cli.StringFlag{Name: `license-description`, Usage: ``},
				cli.StringSliceFlag{Name: `license-required`, Usage: ``},
				cli.StringSliceFlag{Name: `license-permitted`, Usage: ``},
				cli.StringSliceFlag{Name: `license-forbidden`, Usage: ``},
				cli.StringFlag{Name: `license-key`, Usage: ``},
				cli.StringFlag{Name: `license-category`, Usage: ``},
				cli.StringFlag{Name: `license-implementation`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "edit", "edit <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				repository := &github.Repository{
					StargazersCount:  github.Int(c.Int("stargazers-count")),
					UpdatedAt:        &github.Timestamp{now.MustParse(c.String("updated-at"))},
					HasWiki:          github.Bool(c.Bool("has-wiki")),
					CreatedAt:        &github.Timestamp{now.MustParse(c.String("created-at"))},
					Size:             github.Int(c.Int("size")),
					DefaultBranch:    github.String(c.String("default-branch")),
					ForksCount:       github.Int(c.Int("forks-count")),
					PushedAt:         &github.Timestamp{now.MustParse(c.String("pushed-at"))},
					Fork:             github.Bool(c.Bool("fork")),
					SubscribersCount: github.Int(c.Int("subscribers-count")),
					Language:         github.String(c.String("language")),
					WatchersCount:    github.Int(c.Int("watchers-count")),
					Name:             github.String(c.String("name")),
					FullName:         github.String(c.String("full-name")),
					HasDownloads:     github.Bool(c.Bool("has-downloads")),
					NetworkCount:     github.Int(c.Int("network-count")),
					ID:               github.Int(c.Int("id")),
					OpenIssuesCount:  github.Int(c.Int("open-issues-count")),
					Private:          github.Bool(c.Bool("private")),
					HasIssues:        github.Bool(c.Bool("has-issues")),
					Homepage:         github.String(c.String("homepage")),
					Description:      github.String(c.String("description")),
					AutoInit:         github.Bool(c.Bool("auto-init")),
					TeamID:           github.Int(c.Int("team-id")),
					MasterBranch:     github.String(c.String("master-branch")),
				}

				result, res, err := app.gh.Repositories.Edit(owner, repo, repository)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete",
			Usage: `delete a repository.`,
			Description: `delete a repository.

   GitHub API docs: https://developer.github.com/v3/repos/#delete-a-repository`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "delete", "delete <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)

				res, err := app.gh.Repositories.Delete(owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-contributors",
			Usage: `list-contributors lists contributors for a repository.`,
			Description: `list-contributors lists contributors for a repository.

   GitHub API docs: http://developer.github.com/v3/repos/#list-contributors`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `anon`, Usage: `Include anonymous contributors in results or not`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-contributors", "list-contributors <owner> <repository>")
				}

				owner := c.Args().Get(0)
				repository := c.Args().Get(1)
				opt := &github.ListContributorsOptions{
					Anon: c.String("anon"),
				}

				var items []github.Contributor

				for {
					page, res, err := app.gh.Repositories.ListContributors(owner, repository, opt)
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
			Name:  "list-languages",
			Usage: `list-languages lists languages for the specified repository.`,
			Description: `list-languages lists languages for the specified repository. The returned map
   specifies the languages and the number of bytes of code written in that
   language. For example:

       {
         "C": 78769,
         "Python": 7769
       }

   GitHub API Docs: http://developer.github.com/v3/repos/#list-languages`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-languages", "list-languages <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/repos/#list-teams`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-teams", "list-teams <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: https://developer.github.com/v3/repos/#list-tags`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-tags", "list-tags <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/#list-branches`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-branches", "list-branches <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: https://developer.github.com/v3/repos/#get-branch`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-branch", "get-branch <owner> <repo> <branch>")
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

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#list`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-collaborators", "list-collaborators <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#get`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "is-collaborator", "is-collaborator <owner> <repo> <user>")
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

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#add-collaborator`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "add-collaborator", "add-collaborator <owner> <repo> <user>")
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

   GitHub API docs: http://developer.github.com/v3/repos/collaborators/#remove-collaborator`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "remove-collaborator", "remove-collaborator <owner> <repo> <user>")
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

   GitHub API docs: http://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-comments", "list-comments <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-commit-comments", "list-commit-comments <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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
			Usage: `create-comment creates a comment for the given commit.`,
			Description: `create-comment creates a comment for the given commit.
   Note: GitHub allows for comments to be created for non-existing files and positions.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#create-a-commit-comment`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `commit-id`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: `User-mutable fields`},
				cli.StringFlag{Name: `path`, Usage: `User-initialized fields`},
				cli.IntFlag{Name: `position`, Usage: ``},
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
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-comment", "create-comment <owner> <repo> <sha>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				sha := c.Args().Get(2)
				comment := &github.RepositoryComment{
					CommitID:  github.String(c.String("commit-id")),
					ID:        github.Int(c.Int("id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					Body:      github.String(c.String("body")),
					Path:      github.String(c.String("path")),
					Position:  github.Int(c.Int("position")),
				}

				result, res, err := app.gh.Repositories.CreateComment(owner, repo, sha, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-comment",
			Usage: `get-comment gets a single comment from a repository.`,
			Description: `get-comment gets a single comment from a repository.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#get-a-single-commit-comment`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-comment", "get-comment <owner> <repo> <id>")
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
			Usage: `update-comment updates the body of a single comment.`,
			Description: `update-comment updates the body of a single comment.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#update-a-commit-comment`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `body`, Usage: `User-mutable fields`},
				cli.StringFlag{Name: `path`, Usage: `User-initialized fields`},
				cli.IntFlag{Name: `position`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
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
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `commit-id`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "update-comment", "update-comment <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				comment := &github.RepositoryComment{
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					Body:      github.String(c.String("body")),
					Path:      github.String(c.String("path")),
					Position:  github.Int(c.Int("position")),
					ID:        github.Int(c.Int("id")),
					CommitID:  github.String(c.String("commit-id")),
				}

				result, res, err := app.gh.Repositories.UpdateComment(owner, repo, id, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes a single comment from a repository.`,
			Description: `delete-comment deletes a single comment from a repository.

   GitHub API docs: http://developer.github.com/v3/repos/comments/#delete-a-commit-comment`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-comment", "delete-comment <owner> <repo> <id>")
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
			Usage: `list-commits lists the commits of a repository.`,
			Description: `list-commits lists the commits of a repository.

   GitHub API docs: http://developer.github.com/v3/repos/commits/#list`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `until`, Usage: `Until when should Commits be included in the response.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `sha`, Usage: `SHA or branch to start listing Commits from.`},
				cli.StringFlag{Name: `path`, Usage: `Path that should be touched by the returned Commits.`},
				cli.StringFlag{Name: `author`, Usage: `Author of by which to filter Commits.`},
				cli.StringFlag{Name: `since`, Usage: `Since when should Commits be included in the response.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-commits", "list-commits <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.CommitsListOptions{
					SHA:    c.String("sha"),
					Path:   c.String("path"),
					Author: c.String("author"),
					Since:  now.MustParse(c.String("since")),
					Until:  now.MustParse(c.String("until")),
				}

				var items []github.RepositoryCommit

				for {
					page, res, err := app.gh.Repositories.ListCommits(owner, repo, opt)
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
			Name:  "get-commit",
			Usage: `get-commit fetches the specified commit, including all details about it.`,
			Description: `get-commit fetches the specified commit, including all details about it.
   todo: support media formats - https://github.com/google/go-github/issues/6

   GitHub API docs: http://developer.github.com/v3/repos/commits/#get-a-single-commit
   See also: http://developer.github.com//v3/git/commits/#get-a-single-commit provides the same functionality`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-commit", "get-commit <owner> <repo> <sha>")
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
			Usage: `compare-commits compares a range of commits with each other.`,
			Description: `compare-commits compares a range of commits with each other.
   todo: support media formats - https://github.com/google/go-github/issues/6

   GitHub API docs: http://developer.github.com/v3/repos/commits/index.html#compare-two-commits`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					showHelp(c, "compare-commits", "compare-commits <owner> <repo> <base> <head>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				base := c.Args().Get(2)
				head := c.Args().Get(3)

				result, res, err := app.gh.Repositories.CompareCommits(owner, repo, base, head)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-readme",
			Usage: `get-readme gets the Readme file for the repository.`,
			Description: `get-readme gets the Readme file for the repository.

   GitHub API docs: http://developer.github.com/v3/repos/contents/#get-the-readme`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `ref`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-readme", "get-readme <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.RepositoryContentGetOptions{
					Ref: c.String("ref"),
				}

				result, res, err := app.gh.Repositories.GetReadme(owner, repo, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "download-contents",
			Usage: `download-contents returns an io.ReadCloser that reads the contents of the specified file.`,
			Description: `download-contents returns an io.ReadCloser that reads the contents of the
   specified file. This function will work with files of any size, as opposed
   to GetContents which is limited to 1 Mb files. It is the caller's
   responsibility to close the ReadCloser.`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `ref`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "download-contents", "download-contents <owner> <repo> <filepath>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				filepath := c.Args().Get(2)
				opt := &github.RepositoryContentGetOptions{
					Ref: c.String("ref"),
				}

				_, err := app.gh.Repositories.DownloadContents(owner, repo, filepath, opt)
				check(err)

			},
		}, cli.Command{
			Name:  "get-contents",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "create-file",
			Usage: `create-file creates a new file in a repository at the given path and returns the commit and file metadata.`,
			Description: `create-file creates a new file in a repository at the given path and returns
   the commit and file metadata.

   GitHub API docs: http://developer.github.com/v3/repos/contents/#create-a-file`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sha`, Usage: ``},
				cli.StringFlag{Name: `branch`, Usage: ``},
				cli.StringFlag{Name: `author-date`, Usage: ``},
				cli.StringFlag{Name: `author-name`, Usage: ``},
				cli.StringFlag{Name: `author-email`, Usage: ``},
				cli.StringFlag{Name: `committer-date`, Usage: ``},
				cli.StringFlag{Name: `committer-name`, Usage: ``},
				cli.StringFlag{Name: `committer-email`, Usage: ``},
				cli.StringFlag{Name: `message`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-file", "create-file <owner> <repo> <path>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				path := c.Args().Get(2)
				opt := &github.RepositoryContentFileOptions{
					Branch:  github.String(c.String("branch")),
					Message: github.String(c.String("message")),
					SHA:     github.String(c.String("sha")),
				}

				result, res, err := app.gh.Repositories.CreateFile(owner, repo, path, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "update-file",
			Usage: `update-file updates a file in a repository at the given path and returns the commit and file metadata.`,
			Description: `update-file updates a file in a repository at the given path and returns the
   commit and file metadata. Requires the blob SHA of the file being updated.

   GitHub API docs: http://developer.github.com/v3/repos/contents/#update-a-file`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `message`, Usage: ``},
				cli.StringFlag{Name: `sha`, Usage: ``},
				cli.StringFlag{Name: `branch`, Usage: ``},
				cli.StringFlag{Name: `author-date`, Usage: ``},
				cli.StringFlag{Name: `author-name`, Usage: ``},
				cli.StringFlag{Name: `author-email`, Usage: ``},
				cli.StringFlag{Name: `committer-date`, Usage: ``},
				cli.StringFlag{Name: `committer-name`, Usage: ``},
				cli.StringFlag{Name: `committer-email`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "update-file", "update-file <owner> <repo> <path>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				path := c.Args().Get(2)
				opt := &github.RepositoryContentFileOptions{
					Branch:  github.String(c.String("branch")),
					Message: github.String(c.String("message")),
					SHA:     github.String(c.String("sha")),
				}

				result, res, err := app.gh.Repositories.UpdateFile(owner, repo, path, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-file",
			Usage: `delete-file deletes a file from a repository and returns the commit.`,
			Description: `delete-file deletes a file from a repository and returns the commit.
   Requires the blob SHA of the file to be deleted.

   GitHub API docs: http://developer.github.com/v3/repos/contents/#delete-a-file`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `message`, Usage: ``},
				cli.StringFlag{Name: `sha`, Usage: ``},
				cli.StringFlag{Name: `branch`, Usage: ``},
				cli.StringFlag{Name: `author-date`, Usage: ``},
				cli.StringFlag{Name: `author-name`, Usage: ``},
				cli.StringFlag{Name: `author-email`, Usage: ``},
				cli.StringFlag{Name: `committer-name`, Usage: ``},
				cli.StringFlag{Name: `committer-email`, Usage: ``},
				cli.StringFlag{Name: `committer-date`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-file", "delete-file <owner> <repo> <path>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				path := c.Args().Get(2)
				opt := &github.RepositoryContentFileOptions{
					SHA:     github.String(c.String("sha")),
					Branch:  github.String(c.String("branch")),
					Message: github.String(c.String("message")),
				}

				result, res, err := app.gh.Repositories.DeleteFile(owner, repo, path, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-deployments",
			Usage: `list-deployments lists the deployments of a repository.`,
			Description: `list-deployments lists the deployments of a repository.

   GitHub API docs: https://developer.github.com/v3/repos/deployments/#list-deployments`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sha`, Usage: `SHA of the Deployment.`},
				cli.StringFlag{Name: `ref`, Usage: `List deployments for a given ref.`},
				cli.StringFlag{Name: `task`, Usage: `List deployments for a given task.`},
				cli.StringFlag{Name: `environment`, Usage: `List deployments for a given environment.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-deployments", "list-deployments <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.DeploymentsListOptions{
					SHA:         c.String("sha"),
					Ref:         c.String("ref"),
					Task:        c.String("task"),
					Environment: c.String("environment"),
				}

				var items []github.Deployment

				for {
					page, res, err := app.gh.Repositories.ListDeployments(owner, repo, opt)
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
			Name:  "create-deployment",
			Usage: `create-deployment creates a new deployment for a repository.`,
			Description: `create-deployment creates a new deployment for a repository.

   GitHub API docs: https://developer.github.com/v3/repos/deployments/#create-a-deployment`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `required-contexts`, Usage: ``},
				cli.StringFlag{Name: `payload`, Usage: ``},
				cli.StringFlag{Name: `environment`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
				cli.StringFlag{Name: `ref`, Usage: ``},
				cli.StringFlag{Name: `task`, Usage: ``},
				cli.BoolFlag{Name: `auto-merge`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-deployment", "create-deployment <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				request := &github.DeploymentRequest{
					Task:             github.String(c.String("task")),
					AutoMerge:        github.Bool(c.Bool("auto-merge")),
					RequiredContexts: stringSlicePointer(c.StringSlice("required-contexts")),
					Payload:          github.String(c.String("payload")),
					Environment:      github.String(c.String("environment")),
					Description:      github.String(c.String("description")),
					Ref:              github.String(c.String("ref")),
				}

				result, res, err := app.gh.Repositories.CreateDeployment(owner, repo, request)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-deployment-statuses",
			Usage: `list-deployment-statuses lists the statuses of a given deployment of a repository.`,
			Description: `list-deployment-statuses lists the statuses of a given deployment of a repository.

   GitHub API docs: https://developer.github.com/v3/repos/deployments/#list-deployment-statuses`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-deployment-statuses", "list-deployment-statuses <owner> <repo> <deployment>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				deployment, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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
			Usage: `create-deployment-status creates a new status for a deployment.`,
			Description: `create-deployment-status creates a new status for a deployment.

   GitHub API docs: https://developer.github.com/v3/repos/deployments/#create-a-deployment-status`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.StringFlag{Name: `description`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-deployment-status", "create-deployment-status <owner> <repo> <deployment>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				deployment, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				request := &github.DeploymentStatusRequest{
					State:       github.String(c.String("state")),
					Description: github.String(c.String("description")),
				}

				result, res, err := app.gh.Repositories.CreateDeploymentStatus(owner, repo, deployment, request)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-forks",
			Usage: `list-forks lists the forks of the specified repository.`,
			Description: `list-forks lists the forks of the specified repository.

   GitHub API docs: http://developer.github.com/v3/repos/forks/#list-forks`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sort`, Usage: `How to sort the forks list.  Possible values are: newest, oldest,
watchers.  Default is "newest".`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-forks", "list-forks <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.RepositoryListForksOptions{
					Sort: c.String("sort"),
				}

				var items []github.Repository

				for {
					page, res, err := app.gh.Repositories.ListForks(owner, repo, opt)
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
			Name:  "create-fork",
			Usage: `create-fork creates a fork of the specified repository.`,
			Description: `create-fork creates a fork of the specified repository.

   GitHub API docs: http://developer.github.com/v3/repos/forks/#list-forks`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `organization`, Usage: `The organization to fork the repository into.`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-fork", "create-fork <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.RepositoryCreateForkOptions{
					Organization: c.String("organization"),
				}

				result, res, err := app.gh.Repositories.CreateFork(owner, repo, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-hook",
			Usage: `create-hook creates a Hook for the specified repository.`,
			Description: `create-hook creates a Hook for the specified repository.
   Name and Config are required fields.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#create-a-hook`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: `events`, Usage: ``},
				cli.BoolFlag{Name: `active`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-hook", "create-hook <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				hook := &github.Hook{
					ID:        github.Int(c.Int("id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					Name:      github.String(c.String("name")),
					Events:    c.StringSlice("events"),
					Active:    github.Bool(c.Bool("active")),
				}

				result, res, err := app.gh.Repositories.CreateHook(owner, repo, hook)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-hooks",
			Usage: `list-hooks lists all Hooks for the specified repository.`,
			Description: `list-hooks lists all Hooks for the specified repository.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#list`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-hooks", "list-hooks <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#get-single-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-hook", "get-hook <owner> <repo> <id>")
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
			Usage: `edit-hook updates a specified Hook.`,
			Description: `edit-hook updates a specified Hook.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#edit-a-hook`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringSliceFlag{Name: `events`, Usage: ``},
				cli.BoolFlag{Name: `active`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-hook", "edit-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				hook := &github.Hook{
					Name:      github.String(c.String("name")),
					Events:    c.StringSlice("events"),
					Active:    github.Bool(c.Bool("active")),
					ID:        github.Int(c.Int("id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Repositories.EditHook(owner, repo, id, hook)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-hook",
			Usage: `delete-hook deletes a specified Hook.`,
			Description: `delete-hook deletes a specified Hook.

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#delete-a-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-hook", "delete-hook <owner> <repo> <id>")
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

   GitHub API docs: https://developer.github.com/v3/repos/hooks/#ping-a-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "ping-hook", "ping-hook <owner> <repo> <id>")
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

   GitHub API docs: http://developer.github.com/v3/repos/hooks/#test-a-push-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "test-hook", "test-hook <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)

				res, err := app.gh.Repositories.TestHook(owner, repo, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:        "list-service-hooks",
			Usage:       `list-service-hooks is deprecated.`,
			Description: `list-service-hooks is deprecated.  Use Client.list-service-hooks instead.`,
			Flags:       []cli.Flag{},
			Action: func(c *cli.Context) {

				result, res, err := app.gh.Repositories.ListServiceHooks()
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-keys",
			Usage: `list-keys lists the deploy keys for a repository.`,
			Description: `list-keys lists the deploy keys for a repository.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#list`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-keys", "list-keys <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/keys/#get`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-key", "get-key <owner> <repo> <id>")
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
			Usage: `create-key adds a deploy key for a repository.`,
			Description: `create-key adds a deploy key for a repository.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#create`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `key`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-key", "create-key <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				key := &github.Key{
					ID:    github.Int(c.Int("id")),
					Key:   github.String(c.String("key")),
					Title: github.String(c.String("title")),
				}

				result, res, err := app.gh.Repositories.CreateKey(owner, repo, key)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-key",
			Usage: `edit-key edits a deploy key.`,
			Description: `edit-key edits a deploy key.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#edit`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `key`, Usage: ``},
				cli.StringFlag{Name: `title`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-key", "edit-key <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				key := &github.Key{
					Key:   github.String(c.String("key")),
					Title: github.String(c.String("title")),
					ID:    github.Int(c.Int("id")),
				}

				result, res, err := app.gh.Repositories.EditKey(owner, repo, id, key)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-key",
			Usage: `delete-key deletes a deploy key.`,
			Description: `delete-key deletes a deploy key.

   GitHub API docs: http://developer.github.com/v3/repos/keys/#delete`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-key", "delete-key <owner> <repo> <id>")
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
			Usage: `merge a branch in the specified repository.`,
			Description: `merge a branch in the specified repository.

   GitHub API docs: https://developer.github.com/v3/repos/merging/#perform-a-merge`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `commit-message`, Usage: ``},
				cli.StringFlag{Name: `base`, Usage: ``},
				cli.StringFlag{Name: `head`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "merge", "merge <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				request := &github.RepositoryMergeRequest{
					Base:          github.String(c.String("base")),
					Head:          github.String(c.String("head")),
					CommitMessage: github.String(c.String("commit-message")),
				}

				result, res, err := app.gh.Repositories.Merge(owner, repo, request)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-pages-info",
			Usage: `get-pages-info fetches information about a GitHub Pages site.`,
			Description: `get-pages-info fetches information about a GitHub Pages site.

   GitHub API docs: https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-pages-info", "get-pages-info <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/repos/pages/#list-pages-builds`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-pages-builds", "list-pages-builds <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/repos/pages/#list-latest-pages-build`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-latest-pages-build", "get-latest-pages-build <owner> <repo>")
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

   GitHub API docs: http://developer.github.com/v3/repos/releases/#list-releases-for-a-repository`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-releases", "list-releases <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs: http://developer.github.com/v3/repos/releases/#get-a-single-release`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-release", "get-release <owner> <repo> <id>")
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

   GitHub API docs: https://developer.github.com/v3/repos/releases/#get-the-latest-release`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-latest-release", "get-latest-release <owner> <repo>")
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

   GitHub API docs: https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-release-by-tag", "get-release-by-tag <owner> <repo> <tag>")
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
			Usage: `create-release adds a new release for a repository.`,
			Description: `create-release adds a new release for a repository.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#create-a-release`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.BoolFlag{Name: `draft`, Usage: ``},
				cli.StringFlag{Name: `published-at`, Usage: ``},
				cli.StringFlag{Name: `tag-name`, Usage: ``},
				cli.BoolFlag{Name: `prerelease`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `target-commitish`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create-release", "create-release <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				release := &github.RepositoryRelease{
					CreatedAt:       &github.Timestamp{now.MustParse(c.String("created-at"))},
					TargetCommitish: github.String(c.String("target-commitish")),
					Body:            github.String(c.String("body")),
					ID:              github.Int(c.Int("id")),
					Name:            github.String(c.String("name")),
					Draft:           github.Bool(c.Bool("draft")),
					PublishedAt:     &github.Timestamp{now.MustParse(c.String("published-at"))},
					TagName:         github.String(c.String("tag-name")),
					Prerelease:      github.Bool(c.Bool("prerelease")),
				}

				result, res, err := app.gh.Repositories.CreateRelease(owner, repo, release)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-release",
			Usage: `edit-release edits a repository release.`,
			Description: `edit-release edits a repository release.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#edit-a-release`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `target-commitish`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.BoolFlag{Name: `draft`, Usage: ``},
				cli.StringFlag{Name: `published-at`, Usage: ``},
				cli.StringFlag{Name: `tag-name`, Usage: ``},
				cli.BoolFlag{Name: `prerelease`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-release", "edit-release <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				release := &github.RepositoryRelease{
					ID:              github.Int(c.Int("id")),
					Name:            github.String(c.String("name")),
					Draft:           github.Bool(c.Bool("draft")),
					PublishedAt:     &github.Timestamp{now.MustParse(c.String("published-at"))},
					TagName:         github.String(c.String("tag-name")),
					Prerelease:      github.Bool(c.Bool("prerelease")),
					CreatedAt:       &github.Timestamp{now.MustParse(c.String("created-at"))},
					TargetCommitish: github.String(c.String("target-commitish")),
					Body:            github.String(c.String("body")),
				}

				result, res, err := app.gh.Repositories.EditRelease(owner, repo, id, release)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-release",
			Usage: `delete-release delete a single release from a repository.`,
			Description: `delete-release delete a single release from a repository.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#delete-a-release`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-release", "delete-release <owner> <repo> <id>")
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

   GitHub API docs : http://developer.github.com/v3/repos/releases/#list-assets-for-a-release`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-release-assets", "list-release-assets <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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

   GitHub API docs : http://developer.github.com/v3/repos/releases/#get-a-single-release-asset`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-release-asset", "get-release-asset <owner> <repo> <id>")
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
			Usage: `edit-release-asset edits a repository release asset.`,
			Description: `edit-release-asset edits a repository release asset.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#edit-a-release-asset`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `uploader-following`, Usage: ``},
				cli.StringFlag{Name: `uploader-type`, Usage: ``},
				cli.StringFlag{Name: `uploader-login`, Usage: ``},
				cli.StringFlag{Name: `uploader-company`, Usage: ``},
				cli.StringFlag{Name: `uploader-email`, Usage: ``},
				cli.StringFlag{Name: `uploader-name`, Usage: ``},
				cli.IntFlag{Name: `uploader-public-gists`, Usage: ``},
				cli.IntFlag{Name: `uploader-followers`, Usage: ``},
				cli.IntFlag{Name: `uploader-private-gists`, Usage: ``},
				cli.IntFlag{Name: `uploader-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `uploader-collaborators`, Usage: ``},
				cli.StringFlag{Name: `uploader-plan-name`, Usage: ``},
				cli.IntFlag{Name: `uploader-plan-space`, Usage: ``},
				cli.IntFlag{Name: `uploader-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `uploader-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `uploader-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `uploader-blog`, Usage: ``},
				cli.StringFlag{Name: `uploader-created-at`, Usage: ``},
				cli.IntFlag{Name: `uploader-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `uploader-id`, Usage: ``},
				cli.IntFlag{Name: `uploader-public-repos`, Usage: ``},
				cli.StringFlag{Name: `uploader-updated-at`, Usage: ``},
				cli.StringFlag{Name: `uploader-bio`, Usage: ``},
				cli.BoolFlag{Name: `uploader-site-admin`, Usage: ``},
				cli.StringFlag{Name: `uploader-location`, Usage: ``},
				cli.IntFlag{Name: `uploader-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `uploader-hireable`, Usage: ``},
				cli.StringFlag{Name: `label`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.StringFlag{Name: `content-type`, Usage: ``},
				cli.IntFlag{Name: `size`, Usage: ``},
				cli.IntFlag{Name: `download-count`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-release-asset", "edit-release-asset <owner> <repo> <id>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				release := &github.ReleaseAsset{
					Label:         github.String(c.String("label")),
					State:         github.String(c.String("state")),
					ContentType:   github.String(c.String("content-type")),
					Size:          github.Int(c.Int("size")),
					DownloadCount: github.Int(c.Int("download-count")),
					ID:            github.Int(c.Int("id")),
					Name:          github.String(c.String("name")),
					CreatedAt:     &github.Timestamp{now.MustParse(c.String("created-at"))},
					UpdatedAt:     &github.Timestamp{now.MustParse(c.String("updated-at"))},
				}

				result, res, err := app.gh.Repositories.EditReleaseAsset(owner, repo, id, release)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-release-asset",
			Usage: `delete-release-asset delete a single release asset from a repository.`,
			Description: `delete-release-asset delete a single release asset from a repository.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#delete-a-release-asset`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "delete-release-asset", "delete-release-asset <owner> <repo> <id>")
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
			Usage: `upload-release-asset creates an asset by uploading a file into a release repository.`,
			Description: `upload-release-asset creates an asset by uploading a file into a release repository.
   To upload assets that cannot be represented by an os.File, call NewUploadRequest directly.

   GitHub API docs : http://developer.github.com/v3/repos/releases/#upload-a-release-asset`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `name`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 4 {
					showHelp(c, "upload-release-asset", "upload-release-asset <owner> <repo> <id> <file>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				id, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.UploadOptions{
					Name: c.String("name"),
				}
				file, err := os.Open(c.Args().Get(4))
				check(err)

				result, res, err := app.gh.Repositories.UploadReleaseAsset(owner, repo, id, opt, file)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

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

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#contributors`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-contributors-stats", "list-contributors-stats <owner> <repo>")
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

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#commit-activity`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-commit-activity", "list-commit-activity <owner> <repo>")
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

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#code-frequency`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-code-frequency", "list-code-frequency <owner> <repo>")
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

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#participation`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-participation", "list-participation <owner> <repo>")
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

   GitHub API Docs: https://developer.github.com/v3/repos/statistics/#punch-card`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list-punch-card", "list-punch-card <owner> <repo>")
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

   GitHub API docs: http://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-statuses", "list-statuses <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
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
			Usage: `create-status creates a new status for a repository at the specified reference.`,
			Description: `create-status creates a new status for a repository at the specified
   reference.  Ref can be a SHA, a branch name, or a tag name.

   GitHub API docs: http://developer.github.com/v3/repos/statuses/#create-a-status`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `context`, Usage: `A string label to differentiate this status from the statuses of other systems.`},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: `State is the current state of the repository.  Possible values are:
pending, success, error, or failure.`},
				cli.StringFlag{Name: `description`, Usage: `Description is a short high level summary of the status.`},
				cli.StringFlag{Name: `creator-email`, Usage: ``},
				cli.IntFlag{Name: `creator-following`, Usage: ``},
				cli.StringFlag{Name: `creator-type`, Usage: ``},
				cli.StringFlag{Name: `creator-login`, Usage: ``},
				cli.StringFlag{Name: `creator-company`, Usage: ``},
				cli.IntFlag{Name: `creator-private-gists`, Usage: ``},
				cli.StringFlag{Name: `creator-name`, Usage: ``},
				cli.IntFlag{Name: `creator-public-gists`, Usage: ``},
				cli.IntFlag{Name: `creator-followers`, Usage: ``},
				cli.IntFlag{Name: `creator-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `creator-collaborators`, Usage: ``},
				cli.StringFlag{Name: `creator-plan-name`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-space`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `creator-plan-private-repos`, Usage: ``},
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
				cli.StringFlag{Name: `created-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-status", "create-status <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)
				status := &github.RepoStatus{
					ID:          github.Int(c.Int("id")),
					State:       github.String(c.String("state")),
					Description: github.String(c.String("description")),
					CreatedAt:   timePointer(now.MustParse(c.String("created-at"))),
					Context:     github.String(c.String("context")),
					UpdatedAt:   timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Repositories.CreateStatus(owner, repo, ref, status)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "get-combined-status",
			Usage: `get-combined-status returns the combined status of a repository at the specified reference.`,
			Description: `get-combined-status returns the combined status of a repository at the specified
   reference.  ref can be a SHA, a branch name, or a tag name.

   GitHub API docs: https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "get-combined-status", "get-combined-status <owner> <repo> <ref>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				ref := c.Args().Get(2)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				result, res, err := app.gh.Repositories.GetCombinedStatus(owner, repo, ref, opt)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, RepositoriesService)
}
