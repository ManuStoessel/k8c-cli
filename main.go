package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: actionSubCommands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var actionSubCommands = []*cli.Command{
	{
		Name:    "create",
		Aliases: []string{"add"},
		Usage:   "create a new resource",
		Action: func(c *cli.Context) error {
			fmt.Println("added resource: ", c.Args().First())
			return nil
		},
	},
	{
		Name:    "get",
		Aliases: []string{"fetch", "retrieve"},
		Usage:   "retrieve resources",
		Action: func(c *cli.Context) error {
			fmt.Println("fetched resource: ", c.Args().First())
			return nil
		},
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
		Name:  "project",
		Usage: "a project resource",
		Action: func(c *cli.Context) error {
			fmt.Println("project deleted: ", c.Args().First())
			return nil
		},
	},
	{
		Name:  "cluster",
		Usage: "a cluster resource",
		Action: func(c *cli.Context) error {
			fmt.Println("cluster deleted: ", c.Args().First())
			return nil
		},
	},
}
