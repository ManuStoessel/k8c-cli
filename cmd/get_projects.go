/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"strconv"

	"github.com/ManuStoessel/k8c-cli/client"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var getProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Lists projects or fetches a specific named project.",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("projects called")
		k8client, err := client.NewClient("https://run.lab.kubermatic.io", "atokenthatissupersecure")
		if err != nil {
			fmt.Println("Could not initialize Kubermatic API client.")
			return
		}

		projects, err := k8client.ListProjects()
		if err != nil {
			fmt.Println("Error fetching projects.")
			return
		}

		//fmt.Printf("Projects: %+v\n", projects)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleBold)
		t.AppendHeader(table.Row{"ID", "Name", "Clusters", "Status", "Created"})

		rows := make([]table.Row, len(projects))

		for i, p := range projects {
			rows[i] = table.Row{p.ID, p.Name, strconv.FormatInt(p.ClustersNumber, 10), p.Status, p.CreationTimestamp}
		}
		t.AppendRows(rows)

		t.Render()
	},
}

func init() {
	getCmd.AddCommand(getProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
