package main

import (
	"fmt"

	"github.com/ManuStoessel/k8c-cli/client"
	"github.com/urfave/cli/v2"
)

func actionEntry(c *cli.Context) error {

	k8cClient, err := client.NewClient(c.String("url"), c.String("token"))
	if err != nil {
		return err
	}

	projects, err := k8cClient.ListProjects()
	if err != nil {
		return err
	}

	if len(projects) == 0 {
		fmt.Println("No projects found")
		return nil
	}

	for _, project := range projects {
		fmt.Println(project.Name)
	}

	return nil
}
