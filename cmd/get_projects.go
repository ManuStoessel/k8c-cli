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
	"github.com/spf13/cobra"
)

var listAll bool

// projectsCmd represents the projects command
var getProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Lists projects.",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("projects called")
		k8client, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		projects, err := k8client.ListProjects(listAll)
		if err != nil {
			fmt.Println("Error fetching projects.")
			return
		}

		//fmt.Printf("Projects: %+v\n", projects)

		renderProjectList(projects)
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)

	getProjectsCmd.Flags().BoolVarP(&listAll, "listAll", "A", false, "Display all projects the users is allowed to see.")
}
