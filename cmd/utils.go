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
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
	"github.com/kubermatic/go-kubermatic/models"
)

func renderJSON(o interface{}) {
	output, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("Error encoding resources as JSON: %+v\n", o)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", output)
}

func renderProjectList(projects []models.Project) {
	if jsonOutput {
		renderJSON(projects)
	} else {
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
	}
}
