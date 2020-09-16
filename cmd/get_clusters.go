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

var pID string

// clustersCmd represents the clusters command
var getClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Lists clusters for a given project.",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("clusters called")
		k8client, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		clusters, err := k8client.ListClusters(pID)
		if err != nil {
			fmt.Println("Error fetching clusters.")
			return
		}

		//fmt.Printf("%+v", clusters)
		renderClusterList(clusters)
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)

	getClustersCmd.Flags().StringVarP(&pID, "projectID", "p", "", "ID of the project to list clusters for.")
}
