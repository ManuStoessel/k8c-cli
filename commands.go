package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var actionSubCommands = []*cli.Command{
	{
		Name:        "create",
		Aliases:     []string{"add"},
		Usage:       "create a new resource",
		Subcommands: resourceSubCommands,
	},
	{
		Name:        "get",
		Aliases:     []string{"fetch", "list"},
		Usage:       "lists resources of a given type",
		Subcommands: resourceSubCommands,
	},
	{
		Name:        "delete",
		Aliases:     []string{"remove", "rm"},
		Usage:       "delete resources",
		Subcommands: resourceSubCommands,
	},
}

var resourceSubCommands = []*cli.Command{
	{
		Name:  "projects",
		Usage: "project resource type",
		Action: func(c *cli.Context) error {
			fmt.Println("project deleted: ", c.Args().First())
			return nil
		},
	},
	{
		Name:  "clusters",
		Usage: "cluster resource type",
		Action: func(c *cli.Context) error {
			fmt.Println("cluster deleted: ", c.Args().First())
			return nil
		},
	},
}
