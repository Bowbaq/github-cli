package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
	"github.com/kr/pretty"
)

var PullRequestsService = cli.Command{
	Name:     "pull-requests",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list the pull requests for the specified repository.`,
			Description: `list the pull requests for the specified repository.

   GitHub API docs: http://developer.github.com/v3/pulls/#list-pull-requests`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `base`, Usage: `Base filters pull requests by base branch name.`},
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort pull requests. Possible values are: created,
updated, popularity, long-running. Default is "created".`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort pull requests. Possible values are: asc, desc.
If Sort is "created" or not specified, Default is "desc", otherwise Default
is "asc"`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
				cli.StringFlag{Name: `state`, Usage: `State filters pull requests based on their state.  Possible values are:
open, closed.  Default is "open".`},
				cli.StringFlag{Name: `head`, Usage: `Head filters pull requests by head user and branch name in the format of:
"user:ref-name".`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "list", "list <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				opt := &github.PullRequestListOptions{
					State:     c.String("state"),
					Head:      c.String("head"),
					Base:      c.String("base"),
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
				}

				var items []github.PullRequest

				for {
					page, res, err := app.gh.PullRequests.List(owner, repo, opt)
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
			Usage: `get a single pull request.`,
			Description: `get a single pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#get-a-single-pull-request`,
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
			Usage: `create a new pull request on the specified repository.`,
			Description: `create a new pull request on the specified repository.

   GitHub API docs: https://developer.github.com/v3/pulls/#create-a-pull-request`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `base`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.IntFlag{Name: `issue`, Usage: ``},
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.StringFlag{Name: `head`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "create", "create <owner> <repo>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				pull := &github.NewPullRequest{
					Body:  github.String(c.String("body")),
					Issue: github.Int(c.Int("issue")),
					Title: github.String(c.String("title")),
					Head:  github.String(c.String("head")),
					Base:  github.String(c.String("base")),
				}

				result, res, err := app.gh.PullRequests.Create(owner, repo, pull)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit a pull request.`,
			Description: `edit a pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#update-a-pull-request`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.IntFlag{Name: `commits`, Usage: ``},
				cli.StringFlag{Name: `base-repo-default-branch`, Usage: ``},
				cli.IntFlag{Name: `base-repo-size`, Usage: ``},
				cli.StringFlag{Name: `base-repo-pushed-at`, Usage: ``},
				cli.IntFlag{Name: `base-repo-forks-count`, Usage: ``},
				cli.StringFlag{Name: `base-repo-language`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-fork`, Usage: ``},
				cli.IntFlag{Name: `base-repo-subscribers-count`, Usage: ``},
				cli.IntFlag{Name: `base-repo-watchers-count`, Usage: ``},
				cli.StringFlag{Name: `base-repo-name`, Usage: ``},
				cli.StringFlag{Name: `base-repo-full-name`, Usage: ``},
				cli.IntFlag{Name: `base-repo-network-count`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-has-downloads`, Usage: ``},
				cli.IntFlag{Name: `base-repo-id`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-bio`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-owner-site-admin`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-location`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-owner-hireable`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-login`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-company`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-email`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-following`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-type`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-name`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-followers`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-private-gists`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-blog`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-created-at`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-collaborators`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-plan-name`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-plan-space`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-id`, Usage: ``},
				cli.IntFlag{Name: `base-repo-owner-public-repos`, Usage: ``},
				cli.StringFlag{Name: `base-repo-owner-updated-at`, Usage: ``},
				cli.IntFlag{Name: `base-repo-open-issues-count`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-company`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-blog`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-email`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-public-gists`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-updated-at`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-type`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-public-repos`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-following`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-owned-private-repos`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-name`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-followers`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-private-gists`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-disk-usage`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-collaborators`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-login`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-billing-email`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-created-at`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `base-repo-organization-id`, Usage: ``},
				cli.StringFlag{Name: `base-repo-organization-location`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-private`, Usage: `Additional mutable fields when creating and editing a repository`},
				cli.BoolFlag{Name: `base-repo-has-issues`, Usage: ``},
				cli.StringFlag{Name: `base-repo-homepage`, Usage: ``},
				cli.StringFlag{Name: `base-repo-description`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-description`, Usage: ``},
				cli.StringSliceFlag{Name: `base-repo-license-required`, Usage: ``},
				cli.StringSliceFlag{Name: `base-repo-license-permitted`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-body`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-name`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-license-featured`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-category`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-implementation`, Usage: ``},
				cli.StringSliceFlag{Name: `base-repo-license-forbidden`, Usage: ``},
				cli.StringFlag{Name: `base-repo-license-key`, Usage: ``},
				cli.StringFlag{Name: `base-repo-master-branch`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-auto-init`, Usage: ``},
				cli.IntFlag{Name: `base-repo-team-id`, Usage: `Creating an organization repository. Required for non-owners.`},
				cli.IntFlag{Name: `base-repo-stargazers-count`, Usage: ``},
				cli.StringFlag{Name: `base-repo-updated-at`, Usage: ``},
				cli.StringFlag{Name: `base-repo-created-at`, Usage: ``},
				cli.BoolFlag{Name: `base-repo-has-wiki`, Usage: ``},
				cli.StringFlag{Name: `base-user-name`, Usage: ``},
				cli.IntFlag{Name: `base-user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `base-user-followers`, Usage: ``},
				cli.IntFlag{Name: `base-user-private-gists`, Usage: ``},
				cli.IntFlag{Name: `base-user-collaborators`, Usage: ``},
				cli.IntFlag{Name: `base-user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `base-user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `base-user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `base-user-plan-collaborators`, Usage: ``},
				cli.StringFlag{Name: `base-user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `base-user-blog`, Usage: ``},
				cli.StringFlag{Name: `base-user-created-at`, Usage: ``},
				cli.IntFlag{Name: `base-user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `base-user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `base-user-id`, Usage: ``},
				cli.IntFlag{Name: `base-user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `base-user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `base-user-bio`, Usage: ``},
				cli.BoolFlag{Name: `base-user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `base-user-location`, Usage: ``},
				cli.IntFlag{Name: `base-user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `base-user-hireable`, Usage: ``},
				cli.StringFlag{Name: `base-user-type`, Usage: ``},
				cli.StringFlag{Name: `base-user-login`, Usage: ``},
				cli.StringFlag{Name: `base-user-company`, Usage: ``},
				cli.StringFlag{Name: `base-user-email`, Usage: ``},
				cli.IntFlag{Name: `base-user-following`, Usage: ``},
				cli.StringFlag{Name: `base-label`, Usage: ``},
				cli.StringFlag{Name: `base-ref`, Usage: ``},
				cli.StringFlag{Name: `base-sha`, Usage: ``},
				cli.StringFlag{Name: `title`, Usage: ``},
				cli.IntFlag{Name: `comments`, Usage: ``},
				cli.IntFlag{Name: `deletions`, Usage: ``},
				cli.IntFlag{Name: `changed-files`, Usage: ``},
				cli.StringFlag{Name: `head-label`, Usage: ``},
				cli.StringFlag{Name: `head-ref`, Usage: ``},
				cli.StringFlag{Name: `head-sha`, Usage: ``},
				cli.StringFlag{Name: `head-repo-homepage`, Usage: ``},
				cli.StringFlag{Name: `head-repo-description`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-category`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-implementation`, Usage: ``},
				cli.StringSliceFlag{Name: `head-repo-license-forbidden`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-key`, Usage: ``},
				cli.StringSliceFlag{Name: `head-repo-license-required`, Usage: ``},
				cli.StringSliceFlag{Name: `head-repo-license-permitted`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-body`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-name`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-license-featured`, Usage: ``},
				cli.StringFlag{Name: `head-repo-license-description`, Usage: ``},
				cli.StringFlag{Name: `head-repo-master-branch`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-auto-init`, Usage: ``},
				cli.IntFlag{Name: `head-repo-team-id`, Usage: `Creating an organization repository. Required for non-owners.`},
				cli.IntFlag{Name: `head-repo-stargazers-count`, Usage: ``},
				cli.StringFlag{Name: `head-repo-updated-at`, Usage: ``},
				cli.StringFlag{Name: `head-repo-created-at`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-has-wiki`, Usage: ``},
				cli.StringFlag{Name: `head-repo-default-branch`, Usage: ``},
				cli.IntFlag{Name: `head-repo-size`, Usage: ``},
				cli.StringFlag{Name: `head-repo-pushed-at`, Usage: ``},
				cli.IntFlag{Name: `head-repo-forks-count`, Usage: ``},
				cli.StringFlag{Name: `head-repo-language`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-fork`, Usage: ``},
				cli.IntFlag{Name: `head-repo-subscribers-count`, Usage: ``},
				cli.IntFlag{Name: `head-repo-watchers-count`, Usage: ``},
				cli.StringFlag{Name: `head-repo-name`, Usage: ``},
				cli.StringFlag{Name: `head-repo-full-name`, Usage: ``},
				cli.IntFlag{Name: `head-repo-network-count`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-has-downloads`, Usage: ``},
				cli.IntFlag{Name: `head-repo-id`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-collaborators`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-plan-name`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-plan-space`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-blog`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-created-at`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-id`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-public-repos`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-updated-at`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-bio`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-owner-site-admin`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-location`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-owner-hireable`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-following`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-type`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-login`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-company`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-email`, Usage: ``},
				cli.StringFlag{Name: `head-repo-owner-name`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-public-gists`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-followers`, Usage: ``},
				cli.IntFlag{Name: `head-repo-owner-private-gists`, Usage: ``},
				cli.IntFlag{Name: `head-repo-open-issues-count`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-following`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-public-repos`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-name`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-private-gists`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-disk-usage`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-collaborators`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-followers`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-login`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-billing-email`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-created-at`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-location`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-id`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-blog`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-email`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-public-gists`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-updated-at`, Usage: ``},
				cli.IntFlag{Name: `head-repo-organization-total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-type`, Usage: ``},
				cli.StringFlag{Name: `head-repo-organization-company`, Usage: ``},
				cli.BoolFlag{Name: `head-repo-private`, Usage: `Additional mutable fields when creating and editing a repository`},
				cli.BoolFlag{Name: `head-repo-has-issues`, Usage: ``},
				cli.StringFlag{Name: `head-user-name`, Usage: ``},
				cli.IntFlag{Name: `head-user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `head-user-followers`, Usage: ``},
				cli.IntFlag{Name: `head-user-private-gists`, Usage: ``},
				cli.StringFlag{Name: `head-user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `head-user-blog`, Usage: ``},
				cli.StringFlag{Name: `head-user-created-at`, Usage: ``},
				cli.IntFlag{Name: `head-user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-user-collaborators`, Usage: ``},
				cli.StringFlag{Name: `head-user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `head-user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `head-user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `head-user-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `head-user-id`, Usage: ``},
				cli.IntFlag{Name: `head-user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `head-user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `head-user-bio`, Usage: ``},
				cli.BoolFlag{Name: `head-user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `head-user-location`, Usage: ``},
				cli.IntFlag{Name: `head-user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `head-user-hireable`, Usage: ``},
				cli.StringFlag{Name: `head-user-login`, Usage: ``},
				cli.StringFlag{Name: `head-user-company`, Usage: ``},
				cli.StringFlag{Name: `head-user-email`, Usage: ``},
				cli.IntFlag{Name: `head-user-following`, Usage: ``},
				cli.StringFlag{Name: `head-user-type`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: ``},
				cli.StringFlag{Name: `closed-at`, Usage: ``},
				cli.BoolFlag{Name: `merged`, Usage: ``},
				cli.StringFlag{Name: `merged-by-name`, Usage: ``},
				cli.IntFlag{Name: `merged-by-public-gists`, Usage: ``},
				cli.IntFlag{Name: `merged-by-followers`, Usage: ``},
				cli.IntFlag{Name: `merged-by-private-gists`, Usage: ``},
				cli.IntFlag{Name: `merged-by-collaborators`, Usage: ``},
				cli.IntFlag{Name: `merged-by-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `merged-by-plan-name`, Usage: ``},
				cli.IntFlag{Name: `merged-by-plan-space`, Usage: ``},
				cli.IntFlag{Name: `merged-by-plan-collaborators`, Usage: ``},
				cli.StringFlag{Name: `merged-by-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `merged-by-blog`, Usage: ``},
				cli.StringFlag{Name: `merged-by-created-at`, Usage: ``},
				cli.IntFlag{Name: `merged-by-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `merged-by-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `merged-by-id`, Usage: ``},
				cli.IntFlag{Name: `merged-by-public-repos`, Usage: ``},
				cli.StringFlag{Name: `merged-by-updated-at`, Usage: ``},
				cli.StringFlag{Name: `merged-by-bio`, Usage: ``},
				cli.BoolFlag{Name: `merged-by-site-admin`, Usage: ``},
				cli.StringFlag{Name: `merged-by-location`, Usage: ``},
				cli.IntFlag{Name: `merged-by-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `merged-by-hireable`, Usage: ``},
				cli.StringFlag{Name: `merged-by-type`, Usage: ``},
				cli.StringFlag{Name: `merged-by-login`, Usage: ``},
				cli.StringFlag{Name: `merged-by-company`, Usage: ``},
				cli.StringFlag{Name: `merged-by-email`, Usage: ``},
				cli.IntFlag{Name: `merged-by-following`, Usage: ``},
				cli.IntFlag{Name: `additions`, Usage: ``},
				cli.IntFlag{Name: `number`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `merged-at`, Usage: ``},
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
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.BoolFlag{Name: `mergeable`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit", "edit <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				pull := &github.PullRequest{
					Body:         github.String(c.String("body")),
					CreatedAt:    timePointer(now.MustParse(c.String("created-at"))),
					Commits:      github.Int(c.Int("commits")),
					Title:        github.String(c.String("title")),
					Comments:     github.Int(c.Int("comments")),
					Deletions:    github.Int(c.Int("deletions")),
					ChangedFiles: github.Int(c.Int("changed-files")),
					State:        github.String(c.String("state")),
					ClosedAt:     timePointer(now.MustParse(c.String("closed-at"))),
					Merged:       github.Bool(c.Bool("merged")),
					Mergeable:    github.Bool(c.Bool("mergeable")),
					Additions:    github.Int(c.Int("additions")),
					Number:       github.Int(c.Int("number")),
					UpdatedAt:    timePointer(now.MustParse(c.String("updated-at"))),
					MergedAt:     timePointer(now.MustParse(c.String("merged-at"))),
				}

				result, res, err := app.gh.PullRequests.Edit(owner, repo, number, pull)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-commits",
			Usage: `list-commits lists the commits in a pull request.`,
			Description: `list-commits lists the commits in a pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
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
					PerPage: c.Int("per-page"),
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

   GitHub API docs: https://developer.github.com/v3/pulls/#list-pull-requests-files`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
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
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
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

   GitHub API docs: https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged`,
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

   GitHub API docs: https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-buttontrade`,
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
			Usage: `list-comments lists all comments on the specified pull request.`,
			Description: `list-comments lists all comments on the specified pull request.  Specifying a
   pull request number of 0 will return all comments on all pull requests for
   the repository.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `sort`, Usage: `Sort specifies how to sort comments.  Possible values are: created, updated.`},
				cli.StringFlag{Name: `direction`, Usage: `Direction in which to sort comments.  Possible values are: asc, desc.`},
				cli.StringFlag{Name: `since`, Usage: `Since filters comments by time.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "list-comments", "list-comments <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				opt := &github.PullRequestListCommentsOptions{
					Sort:      c.String("sort"),
					Direction: c.String("direction"),
					Since:     now.MustParse(c.String("since")),
				}

				var items []github.PullRequestComment

				for {
					page, res, err := app.gh.PullRequests.ListComments(owner, repo, number, opt)
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
			Usage: `get-comment fetches the specified pull request comment.`,
			Description: `get-comment fetches the specified pull request comment.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#get-a-single-comment`,
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
			Usage: `create-comment creates a new comment on the specified pull request.`,
			Description: `create-comment creates a new comment on the specified pull request.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#create-a-comment`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringFlag{Name: `path`, Usage: ``},
				cli.IntFlag{Name: `position`, Usage: ``},
				cli.StringFlag{Name: `commit-id`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "create-comment", "create-comment <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				comment := &github.PullRequestComment{
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
					Path:      github.String(c.String("path")),
					Position:  github.Int(c.Int("position")),
					CommitID:  github.String(c.String("commit-id")),
				}

				result, res, err := app.gh.PullRequests.CreateComment(owner, repo, number, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-comment",
			Usage: `edit-comment updates a pull request comment.`,
			Description: `edit-comment updates a pull request comment.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#edit-a-comment`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `commit-id`, Usage: ``},
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
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
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
				cli.StringFlag{Name: `body`, Usage: ``},
				cli.StringFlag{Name: `path`, Usage: ``},
				cli.IntFlag{Name: `position`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "edit-comment", "edit-comment <owner> <repo> <number>")
				}

				owner := c.Args().Get(0)
				repo := c.Args().Get(1)
				number, err := strconv.Atoi(c.Args().Get(2))
				check(err)
				comment := &github.PullRequestComment{
					ID:        github.Int(c.Int("id")),
					Body:      github.String(c.String("body")),
					Path:      github.String(c.String("path")),
					Position:  github.Int(c.Int("position")),
					CommitID:  github.String(c.String("commit-id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.PullRequests.EditComment(owner, repo, number, comment)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-comment",
			Usage: `delete-comment deletes a pull request comment.`,
			Description: `delete-comment deletes a pull request comment.

   GitHub API docs: https://developer.github.com/v3/pulls/comments/#delete-a-comment`,
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
