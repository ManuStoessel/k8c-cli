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
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// kubeconfigCmd represents the kubeconfig command
var getKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig [cluster id]",
	Short: "Fetches the kubeconfig for the given cluster",
	Run: func(cmd *cobra.Command, args []string) {
		k8client, err := client.NewClient(baseURL, apiToken)
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		var kubeconfig clientcmdapi.Config

		kubeconfig, err = k8client.GetClusterKubeconfig(pID, seed, args[0])
		if err != nil {
			fmt.Printf("Error fetching kubeconfig: %s\n", err)
		}

		renderKubeconfig(kubeconfig)
	},
}

func init() {
	getCmd.AddCommand(getKubeconfigCmd)

	getKubeconfigCmd.Flags().StringVarP(&pID, "projectID", "p", "", "ID of the project containing the cluster.")
	getKubeconfigCmd.MarkFlagRequired("projectID")

	getKubeconfigCmd.Flags().StringVarP(&seed, "seed", "s", "", "Name of the datacenter containing the cluster.")
	getKubeconfigCmd.MarkFlagRequired("seed")
}
