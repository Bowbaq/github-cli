package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/jinzhu/now"
	"github.com/kr/pretty"
)

var OrganizationsService = cli.Command{
	Name:     "organizations",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name:  "list",
			Usage: `list the organizations for a user.`,
			Description: `list the organizations for a user.  Passing the empty string will list
   organizations for the authenticated user.

   GitHub API docs: http://developer.github.com/v3/orgs/#list-user-organizations`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list", "list <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				var items []github.Organization

				for {
					page, res, err := app.gh.Organizations.List(user, opt)
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
			Usage: `get fetches an organization by name.`,
			Description: `get fetches an organization by name.

   GitHub API docs: http://developer.github.com/v3/orgs/#get-an-organization`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get", "get <org>")
				}

				org := c.Args().Get(0)

				result, res, err := app.gh.Organizations.Get(org)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit",
			Usage: `edit an organization.`,
			Description: `edit an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/#edit-an-organization`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `login`, Usage: ``},
				cli.StringFlag{Name: `billing-email`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `plan-name`, Usage: ``},
				cli.IntFlag{Name: `plan-space`, Usage: ``},
				cli.IntFlag{Name: `plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `location`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.IntFlag{Name: `total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `type`, Usage: ``},
				cli.StringFlag{Name: `company`, Usage: ``},
				cli.StringFlag{Name: `blog`, Usage: ``},
				cli.StringFlag{Name: `email`, Usage: ``},
				cli.IntFlag{Name: `public-gists`, Usage: ``},
				cli.IntFlag{Name: `public-repos`, Usage: ``},
				cli.IntFlag{Name: `following`, Usage: ``},
				cli.IntFlag{Name: `owned-private-repos`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.IntFlag{Name: `followers`, Usage: ``},
				cli.IntFlag{Name: `private-gists`, Usage: ``},
				cli.IntFlag{Name: `disk-usage`, Usage: ``},
				cli.IntFlag{Name: `collaborators`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "edit", "edit <name>")
				}

				name := c.Args().Get(0)
				org := &github.Organization{
					Type:              github.String(c.String("type")),
					Company:           github.String(c.String("company")),
					Blog:              github.String(c.String("blog")),
					Email:             github.String(c.String("email")),
					PublicGists:       github.Int(c.Int("public-gists")),
					UpdatedAt:         timePointer(now.MustParse(c.String("updated-at"))),
					TotalPrivateRepos: github.Int(c.Int("total-private-repos")),
					PublicRepos:       github.Int(c.Int("public-repos")),
					Following:         github.Int(c.Int("following")),
					OwnedPrivateRepos: github.Int(c.Int("owned-private-repos")),
					Name:              github.String(c.String("name")),
					Followers:         github.Int(c.Int("followers")),
					PrivateGists:      github.Int(c.Int("private-gists")),
					DiskUsage:         github.Int(c.Int("disk-usage")),
					Collaborators:     github.Int(c.Int("collaborators")),
					Login:             github.String(c.String("login")),
					BillingEmail:      github.String(c.String("billing-email")),
					CreatedAt:         timePointer(now.MustParse(c.String("created-at"))),
					ID:                github.Int(c.Int("id")),
					Location:          github.String(c.String("location")),
				}

				result, res, err := app.gh.Organizations.Edit(name, org)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-hooks",
			Usage: `list-hooks lists all Hooks for the specified organization.`,
			Description: `list-hooks lists all Hooks for the specified organization.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#list-hooks`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-hooks", "list-hooks <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				var items []github.Hook

				for {
					page, res, err := app.gh.Organizations.ListHooks(org, opt)
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

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#get-single-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-hook", "get-hook <org> <id>")
				}

				org := c.Args().Get(0)
				id, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				result, res, err := app.gh.Organizations.GetHook(org, id)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-hook",
			Usage: `create-hook creates a Hook for the specified org.`,
			Description: `create-hook creates a Hook for the specified org.
   Name and Config are required fields.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#create-a-hook`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `created-at`, Usage: ``},
				cli.StringFlag{Name: `updated-at`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringSliceFlag{Name: `events`, Usage: ``},
				cli.BoolFlag{Name: `active`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "create-hook", "create-hook <org>")
				}

				org := c.Args().Get(0)
				hook := &github.Hook{
					Name:      github.String(c.String("name")),
					Events:    c.StringSlice("events"),
					Active:    github.Bool(c.Bool("active")),
					ID:        github.Int(c.Int("id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
				}

				result, res, err := app.gh.Organizations.CreateHook(org, hook)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-hook",
			Usage: `edit-hook updates a specified Hook.`,
			Description: `edit-hook updates a specified Hook.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#edit-a-hook`,
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
					showHelp(c, "edit-hook", "edit-hook <org> <id>")
				}

				org := c.Args().Get(0)
				id, err := strconv.Atoi(c.Args().Get(1))
				check(err)
				hook := &github.Hook{
					ID:        github.Int(c.Int("id")),
					CreatedAt: timePointer(now.MustParse(c.String("created-at"))),
					UpdatedAt: timePointer(now.MustParse(c.String("updated-at"))),
					Name:      github.String(c.String("name")),
					Events:    c.StringSlice("events"),
					Active:    github.Bool(c.Bool("active")),
				}

				result, res, err := app.gh.Organizations.EditHook(org, id, hook)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "ping-hook",
			Usage: `ping-hook triggers a 'ping' event to be sent to the Hook.`,
			Description: `ping-hook triggers a 'ping' event to be sent to the Hook.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#ping-a-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "ping-hook", "ping-hook <org> <id>")
				}

				org := c.Args().Get(0)
				id, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				res, err := app.gh.Organizations.PingHook(org, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "delete-hook",
			Usage: `delete-hook deletes a specified Hook.`,
			Description: `delete-hook deletes a specified Hook.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#delete-a-hook`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "delete-hook", "delete-hook <org> <id>")
				}

				org := c.Args().Get(0)
				id, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				res, err := app.gh.Organizations.DeleteHook(org, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-members",
			Usage: `list-members lists the members for an organization.`,
			Description: `list-members lists the members for an organization.  If the authenticated
   user is an owner of the organization, this will return both concealed and
   public members, otherwise it will only return public members.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#members-list`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: `public-only`, Usage: `If true (or if the authenticated user is not an owner of the
organization), list only publicly visible members.`},
				cli.StringFlag{Name: `filter`, Usage: `Filter members returned in the list.  Possible values are:
2fa_disabled, all.  Default is "all".`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-members", "list-members <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListMembersOptions{
					PublicOnly: c.Bool("public-only"),
					Filter:     c.String("filter"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Organizations.ListMembers(org, opt)
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
			Name:  "is-member",
			Usage: `is-member checks if a user is a member of an organization.`,
			Description: `is-member checks if a user is a member of an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#check-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "is-member", "is-member <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				result, res, err := app.gh.Organizations.IsMember(org, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "is-public-member",
			Usage: `is-public-member checks if a user is a public member of an organization.`,
			Description: `is-public-member checks if a user is a public member of an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#check-public-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "is-public-member", "is-public-member <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				result, res, err := app.gh.Organizations.IsPublicMember(org, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "remove-member",
			Usage: `remove-member removes a user from all teams of an organization.`,
			Description: `remove-member removes a user from all teams of an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#remove-a-member`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "remove-member", "remove-member <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				res, err := app.gh.Organizations.RemoveMember(org, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "publicize-membership",
			Usage: `publicize-membership publicizes a user's membership in an organization.`,
			Description: `publicize-membership publicizes a user's membership in an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#publicize-a-users-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "publicize-membership", "publicize-membership <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				res, err := app.gh.Organizations.PublicizeMembership(org, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "conceal-membership",
			Usage: `conceal-membership conceals a user's membership in an organization.`,
			Description: `conceal-membership conceals a user's membership in an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#conceal-a-users-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "conceal-membership", "conceal-membership <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				res, err := app.gh.Organizations.ConcealMembership(org, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-org-memberships",
			Usage: `list-org-memberships lists the organization memberships for the authenticated user.`,
			Description: `list-org-memberships lists the organization memberships for the authenticated user.

   GitHub API docs: https://developer.github.com/v3/orgs/members/#list-your-organization-memberships`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `state`, Usage: `Filter memberships to include only those withe the specified state.
Possible values are: "active", "pending".`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOrgMembershipsOptions{
					State: c.String("state"),
				}

				var items []github.Membership

				for {
					page, res, err := app.gh.Organizations.ListOrgMemberships(opt)
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
			Name:  "get-org-membership",
			Usage: `get-org-membership gets the membership for the authenticated user for the specified organization.`,
			Description: `get-org-membership gets the membership for the authenticated user for the
   specified organization.

   GitHub API docs: https://developer.github.com/v3/orgs/members/#get-your-organization-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get-org-membership", "get-org-membership <org>")
				}

				org := c.Args().Get(0)

				result, res, err := app.gh.Organizations.GetOrgMembership(org)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-org-membership",
			Usage: `edit-org-membership edits the membership for the authenticated user for the specified organization.`,
			Description: `edit-org-membership edits the membership for the authenticated user for the
   specified organization.

   GitHub API docs: https://developer.github.com/v3/orgs/members/#edit-your-organization-membership`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `organization-created-at`, Usage: ``},
				cli.StringFlag{Name: `organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-private-repos`, Usage: ``},
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
				cli.StringFlag{Name: `user-bio`, Usage: ``},
				cli.BoolFlag{Name: `user-site-admin`, Usage: ``},
				cli.IntFlag{Name: `user-disk-usage`, Usage: ``},
				cli.StringFlag{Name: `user-location`, Usage: ``},
				cli.BoolFlag{Name: `user-hireable`, Usage: ``},
				cli.StringFlag{Name: `user-company`, Usage: ``},
				cli.StringFlag{Name: `user-email`, Usage: ``},
				cli.IntFlag{Name: `user-following`, Usage: ``},
				cli.StringFlag{Name: `user-type`, Usage: ``},
				cli.StringFlag{Name: `user-login`, Usage: ``},
				cli.IntFlag{Name: `user-followers`, Usage: ``},
				cli.IntFlag{Name: `user-private-gists`, Usage: ``},
				cli.StringFlag{Name: `user-name`, Usage: ``},
				cli.IntFlag{Name: `user-public-gists`, Usage: ``},
				cli.StringFlag{Name: `user-created-at`, Usage: ``},
				cli.IntFlag{Name: `user-total-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-owned-private-repos`, Usage: ``},
				cli.IntFlag{Name: `user-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `user-plan-private-repos`, Usage: ``},
				cli.StringFlag{Name: `user-plan-name`, Usage: ``},
				cli.IntFlag{Name: `user-plan-space`, Usage: ``},
				cli.StringFlag{Name: `user-gravatar-id`, Usage: ``},
				cli.StringFlag{Name: `user-blog`, Usage: ``},
				cli.StringFlag{Name: `user-updated-at`, Usage: ``},
				cli.IntFlag{Name: `user-id`, Usage: ``},
				cli.IntFlag{Name: `user-public-repos`, Usage: ``},
				cli.StringFlag{Name: `state`, Usage: `State is the user's status within the organization or team.
Possible values are: "active", "pending"`},
				cli.StringFlag{Name: `role`, Usage: `TODO(willnorris): add docs`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "edit-org-membership", "edit-org-membership <org>")
				}

				org := c.Args().Get(0)
				membership := &github.Membership{
					State: github.String(c.String("state")),
					Role:  github.String(c.String("role")),
				}

				result, res, err := app.gh.Organizations.EditOrgMembership(org, membership)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-teams",
			Usage: `list-teams lists all of the teams for an organization.`,
			Description: `list-teams lists all of the teams for an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-teams`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-teams", "list-teams <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				var items []github.Team

				for {
					page, res, err := app.gh.Organizations.ListTeams(org, opt)
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
			Name:  "get-team",
			Usage: `get-team fetches a team by ID.`,
			Description: `get-team fetches a team by ID.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "get-team", "get-team <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				result, res, err := app.gh.Organizations.GetTeam(team)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-team",
			Usage: `create-team creates a new team within an organization.`,
			Description: `create-team creates a new team within an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#create-team`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `repos-count`, Usage: ``},
				cli.IntFlag{Name: `organization-public-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-following`, Usage: ``},
				cli.IntFlag{Name: `organization-owned-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-name`, Usage: ``},
				cli.IntFlag{Name: `organization-followers`, Usage: ``},
				cli.IntFlag{Name: `organization-private-gists`, Usage: ``},
				cli.IntFlag{Name: `organization-disk-usage`, Usage: ``},
				cli.IntFlag{Name: `organization-collaborators`, Usage: ``},
				cli.StringFlag{Name: `organization-login`, Usage: ``},
				cli.StringFlag{Name: `organization-billing-email`, Usage: ``},
				cli.StringFlag{Name: `organization-created-at`, Usage: ``},
				cli.StringFlag{Name: `organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-private-repos`, Usage: ``},
				cli.IntFlag{Name: `organization-id`, Usage: ``},
				cli.StringFlag{Name: `organization-location`, Usage: ``},
				cli.StringFlag{Name: `organization-company`, Usage: ``},
				cli.StringFlag{Name: `organization-blog`, Usage: ``},
				cli.StringFlag{Name: `organization-email`, Usage: ``},
				cli.IntFlag{Name: `organization-public-gists`, Usage: ``},
				cli.StringFlag{Name: `organization-updated-at`, Usage: ``},
				cli.IntFlag{Name: `organization-total-private-repos`, Usage: ``},
				cli.StringFlag{Name: `organization-type`, Usage: ``},
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
				cli.StringFlag{Name: `slug`, Usage: ``},
				cli.StringFlag{Name: `permission`, Usage: ``},
				cli.IntFlag{Name: `members-count`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "create-team", "create-team <org>")
				}

				org := c.Args().Get(0)
				team := &github.Team{
					Name:         github.String(c.String("name")),
					Slug:         github.String(c.String("slug")),
					Permission:   github.String(c.String("permission")),
					MembersCount: github.Int(c.Int("members-count")),
					ReposCount:   github.Int(c.Int("repos-count")),
					ID:           github.Int(c.Int("id")),
				}

				result, res, err := app.gh.Organizations.CreateTeam(org, team)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-team",
			Usage: `edit-team edits a team.`,
			Description: `edit-team edits a team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#edit-team`,
			Flags: []cli.Flag{
				cli.StringFlag{Name: `slug`, Usage: ``},
				cli.StringFlag{Name: `permission`, Usage: ``},
				cli.IntFlag{Name: `members-count`, Usage: ``},
				cli.IntFlag{Name: `repos-count`, Usage: ``},
				cli.StringFlag{Name: `organization-created-at`, Usage: ``},
				cli.StringFlag{Name: `organization-plan-name`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-space`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-collaborators`, Usage: ``},
				cli.IntFlag{Name: `organization-plan-private-repos`, Usage: ``},
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
				cli.IntFlag{Name: `id`, Usage: ``},
				cli.StringFlag{Name: `name`, Usage: ``},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "edit-team", "edit-team <id>")
				}

				id, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				team := &github.Team{
					MembersCount: github.Int(c.Int("members-count")),
					ReposCount:   github.Int(c.Int("repos-count")),
					ID:           github.Int(c.Int("id")),
					Name:         github.String(c.String("name")),
					Slug:         github.String(c.String("slug")),
					Permission:   github.String(c.String("permission")),
				}

				result, res, err := app.gh.Organizations.EditTeam(id, team)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "delete-team",
			Usage: `delete-team deletes a team.`,
			Description: `delete-team deletes a team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#delete-team`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "delete-team", "delete-team <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				res, err := app.gh.Organizations.DeleteTeam(team)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-team-members",
			Usage: `list-team-members lists all of the users who are members of the specified team.`,
			Description: `list-team-members lists all of the users who are members of the specified
   team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-members`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-team-members", "list-team-members <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				var items []github.User

				for {
					page, res, err := app.gh.Organizations.ListTeamMembers(team, opt)
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
			Name:  "is-team-member",
			Usage: `is-team-member checks if a user is a member of the specified team.`,
			Description: `is-team-member checks if a user is a member of the specified team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-member`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "is-team-member", "is-team-member <team> <user>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				user := c.Args().Get(1)

				result, res, err := app.gh.Organizations.IsTeamMember(team, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "list-team-repos",
			Usage: `list-team-repos lists the repositories that the specified team has access to.`,
			Description: `list-team-repos lists the repositories that the specified team has access to.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-repos`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					showHelp(c, "list-team-repos", "list-team-repos <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("per-page"),
				}

				var items []github.Repository

				for {
					page, res, err := app.gh.Organizations.ListTeamRepos(team, opt)
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
			Name:  "is-team-repo",
			Usage: `is-team-repo checks if a team manages the specified repository.`,
			Description: `is-team-repo checks if a team manages the specified repository.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-repo`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "is-team-repo", "is-team-repo <team> <owner> <repo>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				owner := c.Args().Get(1)
				repo := c.Args().Get(2)

				result, res, err := app.gh.Organizations.IsTeamRepo(team, owner, repo)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "add-team-repo",
			Usage: `add-team-repo adds a repository to be managed by the specified team.`,
			Description: `add-team-repo adds a repository to be managed by the specified team.  The
   specified repository must be owned by the organization to which the team
   belongs, or a direct fork of a repository owned by the organization.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#add-team-repo`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "add-team-repo", "add-team-repo <team> <owner> <repo>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				owner := c.Args().Get(1)
				repo := c.Args().Get(2)

				res, err := app.gh.Organizations.AddTeamRepo(team, owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "remove-team-repo",
			Usage: `remove-team-repo removes a repository from being managed by the specified team.`,
			Description: `remove-team-repo removes a repository from being managed by the specified
   team.  Note that this does not delete the repository, it just removes it
   from the team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#remove-team-repo`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					showHelp(c, "remove-team-repo", "remove-team-repo <team> <owner> <repo>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				owner := c.Args().Get(1)
				repo := c.Args().Get(2)

				res, err := app.gh.Organizations.RemoveTeamRepo(team, owner, repo)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-user-teams",
			Usage: `list-user-teams lists a user's teams GitHub API docs: https://developer.github.com/v3/orgs/teams/#list-user-teams`,
			Description: `list-user-teams lists a user's teams
   GitHub API docs: https://developer.github.com/v3/orgs/teams/#list-user-teams`,
			Flags: []cli.Flag{
				cli.IntFlag{Name: `page`, Usage: `For paginated result sets, page of results to retrieve.`},
				cli.IntFlag{Name: `per-page`, Usage: `For paginated result sets, the number of results to include per page.`},
				cli.BoolFlag{Name: `all`, Usage: `For paginated result sets, fetch all remaining pages starting at "page"`},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOptions{
					PerPage: c.Int("per-page"),
					Page:    c.Int("page"),
				}

				var items []github.Team

				for {
					page, res, err := app.gh.Organizations.ListUserTeams(opt)
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
			Name:  "get-team-membership",
			Usage: `get-team-membership returns the membership status for a user in a team.`,
			Description: `get-team-membership returns the membership status for a user in a team.

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#get-team-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "get-team-membership", "get-team-membership <team> <user>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				user := c.Args().Get(1)

				result, res, err := app.gh.Organizations.GetTeamMembership(team, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "add-team-membership",
			Usage: `add-team-membership adds or invites a user to a team.`,
			Description: `add-team-membership adds or invites a user to a team.

   In order to add a membership between a user and a team, the authenticated
   user must have 'admin' permissions to the team or be an owner of the
   organization that the team is associated with.

   If the user is already a part of the team's organization (meaning they're on
   at least one other team in the organization), this endpoint will add the
   user to the team.

   If the user is completely unaffiliated with the team's organization (meaning
   they're on none of the organization's teams), this endpoint will send an
   invitation to the user via email. This newly-created membership will be in
   the "pending" state until the user accepts the invitation, at which point
   the membership will transition to the "active" state and the user will be
   added as a member of the team.

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#add-team-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "add-team-membership", "add-team-membership <team> <user>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				user := c.Args().Get(1)

				result, res, err := app.gh.Organizations.AddTeamMembership(team, user)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "remove-team-membership",
			Usage: `remove-team-membership removes a user from a team.`,
			Description: `remove-team-membership removes a user from a team.

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#remove-team-membership`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					showHelp(c, "remove-team-membership", "remove-team-membership <team> <user>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				user := c.Args().Get(1)

				res, err := app.gh.Organizations.RemoveTeamMembership(team, user)
				checkResponse(res.Response, err)

			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, OrganizationsService)
}
