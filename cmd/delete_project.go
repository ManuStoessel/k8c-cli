/*
Copyright © 2020 Manuel Stößel <manuel@soessel.dev>

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

// projectCmd represents the project command
var deleteProjectCmd = &cobra.Command{
	Use:   "project [id]",
	Short: "Lets you delete a project",
	Run: func(cmd *cobra.Command, args []string) {
		k8client, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		err = k8client.DeleteProject(args[0])
		if err != nil {
			fmt.Printf("Error deleting project %s.\n", args[0])
			return
		}

		fmt.Printf("Successfully deleted project with ID %s", args[0])
	},
}

func init() {
	deleteCmd.AddCommand(deleteProjectCmd)
}
