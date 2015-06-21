package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var RepositoriesService = cli.Command{
	Name:     "repositories",
	HideHelp: true,
	Action:   fixHelp,
	Subcommands: []cli.Command{
		cli.Command{
			Name: "list",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-by-org",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-all",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-contributors",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-languages",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-teams",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-tags",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-branches",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-branch",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-collaborators",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "is-collaborator",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "add-collaborator",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "remove-collaborator",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-comments",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-commit-comments",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "update-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-comment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-commits",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-commit",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "compare-commits",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-readme",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "download-contents",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-contents",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-file",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "update-file",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-file",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-archive-link",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-deployments",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-deployment",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-deployment-statuses",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-deployment-status",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-forks",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-fork",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-hooks",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "ping-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "test-hook",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-service-hooks",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-keys",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-key",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-key",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-key",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-key",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "merge",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-pages-info",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-pages-builds",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-latest-pages-build",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-releases",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-latest-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-release-by-tag",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-single-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-release",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-release-assets",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-release-asset",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "edit-release-asset",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "delete-release-asset",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "upload-release-asset",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-contributors-stats",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-commit-activity",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-code-frequency",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-participation",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-punch-card",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "list-statuses",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "create-status",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		}, cli.Command{
			Name: "get-combined-status",
			Action: func(c *cli.Context) {
				fmt.Println("Not implemented")
				os.Exit(1)
			},
		},
	},
}

func init() {
	app.cli.Commands = append(app.cli.Commands, RepositoriesService)
}
