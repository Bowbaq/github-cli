package main

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
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

   GitHub API docs: http://developer.github.com/v3/orgs/#list-user-organizations
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list <user>")
				}

				user := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: http://developer.github.com/v3/orgs/#get-an-organization
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get <org>")
				}

				org := c.Args().Get(0)

				result, res, err := app.gh.Organizations.Get(org)
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
			Name:  "list-hooks",
			Usage: `list-hooks lists all Hooks for the specified organization.`,
			Description: `list-hooks lists all Hooks for the specified organization.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#list-hooks
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-hooks <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#get-single-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-hook <org> <id>")
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
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-hook",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "ping-hook",
			Usage: `ping-hook triggers a 'ping' event to be sent to the Hook.`,
			Description: `ping-hook triggers a 'ping' event to be sent to the Hook.

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#ping-a-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "ping-hook <org> <id>")
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

   GitHub API docs: https://developer.github.com/v3/orgs/hooks/#delete-a-hook
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "delete-hook <org> <id>")
				}

				org := c.Args().Get(0)
				id, err := strconv.Atoi(c.Args().Get(1))
				check(err)

				res, err := app.gh.Organizations.DeleteHook(org, id)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-members",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "is-member",
			Usage: `is-member checks if a user is a member of an organization.`,
			Description: `is-member checks if a user is a member of an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/members/#check-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "is-member <org> <user>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/members/#check-public-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "is-public-member <org> <user>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/members/#remove-a-member
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "remove-member <org> <user>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/members/#publicize-a-users-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "publicize-membership <org> <user>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/members/#conceal-a-users-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "conceal-membership <org> <user>")
				}

				org := c.Args().Get(0)
				user := c.Args().Get(1)

				res, err := app.gh.Organizations.ConcealMembership(org, user)
				checkResponse(res.Response, err)

			},
		}, cli.Command{
			Name:  "list-org-memberships",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "get-org-membership",
			Usage: `get-org-membership gets the membership for the authenticated user for the specified organization.`,
			Description: `get-org-membership gets the membership for the authenticated user for the
   specified organization.

   GitHub API docs: https://developer.github.com/v3/orgs/members/#get-your-organization-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get-org-membership <org>")
				}

				org := c.Args().Get(0)

				result, res, err := app.gh.Organizations.GetOrgMembership(org)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "edit-org-membership",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "list-teams",
			Usage: `list-teams lists all of the teams for an organization.`,
			Description: `list-teams lists all of the teams for an organization.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-teams
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-teams <org>")
				}

				org := c.Args().Get(0)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "get-team <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)

				result, res, err := app.gh.Organizations.GetTeam(team)
				checkResponse(res.Response, err)
				fmt.Printf("%# v", pretty.Formatter(result))

			},
		}, cli.Command{
			Name:  "create-team",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "edit-team",
			Usage: "not implemented",
			Action: func(c *cli.Context) {
				fatalln("Not implemented")
			},
		}, cli.Command{
			Name:  "delete-team",
			Usage: `delete-team deletes a team.`,
			Description: `delete-team deletes a team.

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#delete-team
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "delete-team <team>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-members
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-team-members <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-member
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "is-team-member <team> <user>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-repos
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
					fatalln("Usage: " + c.App.Name + "list-team-repos <team>")
				}

				team, err := strconv.Atoi(c.Args().Get(0))
				check(err)
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-repo
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "is-team-repo <team> <owner> <repo>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#add-team-repo
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "add-team-repo <team> <owner> <repo>")
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

   GitHub API docs: http://developer.github.com/v3/orgs/teams/#remove-team-repo
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 3 {
					fatalln("Usage: " + c.App.Name + "remove-team-repo <team> <owner> <repo>")
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
   GitHub API docs: https://developer.github.com/v3/orgs/teams/#list-user-teams
`,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "fetch all the pages"},
				cli.IntFlag{Name: "page, p", Value: 0, Usage: "fetch this specific page"},
				cli.IntFlag{Name: "page-size, ps", Value: 30, Usage: "fetch <page-size> items per page"},
			},
			Action: func(c *cli.Context) {
				opt := &github.ListOptions{
					Page:    c.Int("page"),
					PerPage: c.Int("page-size"),
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

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#get-team-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "get-team-membership <team> <user>")
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

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#add-team-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "add-team-membership <team> <user>")
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

   GitHub API docs: https://developer.github.com/v3/orgs/teams/#remove-team-membership
`,
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) {
				if len(c.Args()) < 2 {
					fatalln("Usage: " + c.App.Name + "remove-team-membership <team> <user>")
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
