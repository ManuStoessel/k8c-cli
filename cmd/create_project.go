/*
Copyright © 2020 Manuel Stößel <manuel@stoessel.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/ManuStoessel/k8c-cli/client"
	"github.com/kubermatic/go-kubermatic/models"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var createProjectCmd = &cobra.Command{
	Use:   "project [name]",
	Short: "Lets you create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		k8client, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		//TODO: find a nice way to create a labeled project
		project, err := k8client.CreateProject(args[0], nil)
		if err != nil {
			fmt.Println("Error fetching projects.")
			return
		}

		projects := make([]models.Project, 1)
		projects[0] = project
		renderProjectList(projects)
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
