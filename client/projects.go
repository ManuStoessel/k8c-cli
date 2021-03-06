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

package client

import (
	"errors"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	projectPath       string = ".." + apiV1Path + "/projects"
	clustersSubPath   string = "/clusters"
	datacenterSubPath string = "/dc"
	kubeconfigSubPath string = "/kubeconfig"
)

// ListProjects lists all projects a user has permission to see
// if listall is true, all projects the user has access to, will be listed
// if listall is false (default), only clusters owned by the user will be listed
func (c *Client) ListProjects(listall bool) ([]models.Project, error) {
	req, err := c.newRequest("GET", projectPath, nil)
	if err != nil {
		return nil, err
	}

	if listall {
		params := req.URL.Query()
		params.Add("displayAll", "true")
		req.URL.RawQuery = params.Encode()
	}

	result := make([]models.Project, 0)

	resp, err := c.do(req, &result)
	if err != nil {
		return nil, err
	}

	// StatusCodes 401 and 409 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 409 {
		return nil, nil
	}

	if resp.StatusCode >= 299 {
		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// CreateProject creates a new project with the given name and labels
func (c *Client) CreateProject(name string, labels map[string]string) (models.Project, error) {
	newProject := models.Project{
		Name:   name,
		Labels: labels,
	}

	req, err := c.newRequest("POST", projectPath, newProject)
	if err != nil {
		return models.Project{}, err
	}

	result := models.Project{}

	resp, err := c.do(req, &result)
	if err != nil {
		return models.Project{}, err
	}

	// StatusCodes 401 and 409 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 409 {
		return models.Project{}, nil
	}

	if resp.StatusCode >= 299 {
		return models.Project{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// DeleteProject deletes a project identified by id
func (c *Client) DeleteProject(projectID string) error {
	req, err := c.newRequest("DELETE", projectPath+"/"+projectID, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req, nil)
	if err != nil {
		return err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil
	}

	if resp.StatusCode >= 299 {
		return errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

// ListClustersForProject lists all clusters for a given Project (identified by ID)
func (c *Client) ListClustersForProject(projectID string) ([]models.Cluster, error) {
	req, err := c.newRequest("GET", projectPath+"/"+projectID+clustersSubPath, nil)
	if err != nil {
		return nil, err
	}

	result := make([]models.Cluster, 0)

	resp, err := c.do(req, &result)
	if err != nil {
		return nil, err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil, nil
	}

	if resp.StatusCode >= 299 {
		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// ListClustersForProjectAndDatacenter lists all clusters for a given Project
// (identified by ID) and Seed (identified by name)
func (c *Client) ListClustersForProjectAndDatacenter(projectID string, seed string) ([]models.Cluster, error) {
	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath, nil)
	if err != nil {
		return nil, err
	}

	result := make([]models.Cluster, 0)

	resp, err := c.do(req, &result)
	if err != nil {
		return nil, err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil, nil
	}

	if resp.StatusCode >= 299 {
		return nil, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// GetCluster fetches cluster
func (c *Client) GetCluster(projectID string, seed string, clusterID string) (models.Cluster, error) {
	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID, nil)
	if err != nil {
		return models.Cluster{}, err
	}

	result := models.Cluster{}

	resp, err := c.do(req, &result)
	if err != nil {
		return models.Cluster{}, err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return models.Cluster{}, nil
	}

	if resp.StatusCode >= 299 {
		return models.Cluster{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}

// DeleteCluster deletes a given cluster
func (c *Client) DeleteCluster(projectID string, seed string, clusterID string) error {
	req, err := c.newRequest("DELETE", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req, nil)
	if err != nil {
		return err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil
	}

	if resp.StatusCode >= 299 {
		return errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

// GetClusterKubeconfig fetches kubeconfig for a given cluster
func (c *Client) GetClusterKubeconfig(projectID string, seed string, clusterID string) (clientcmdapi.Config, error) {
	req, err := c.newRequest("GET", projectPath+"/"+projectID+datacenterSubPath+"/"+seed+clustersSubPath+"/"+clusterID+kubeconfigSubPath, nil)
	if err != nil {
		return clientcmdapi.Config{}, err
	}

	result := clientcmdapi.Config{}

	resp, err := c.do(req, &result)
	if err != nil {
		return clientcmdapi.Config{}, err
	}

	// StatusCodes 401 and 403 mean empty response and should be treated as such
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return clientcmdapi.Config{}, nil
	}

	if resp.StatusCode >= 299 {
		return clientcmdapi.Config{}, errors.New("Got non-2xx return code: " + strconv.Itoa(resp.StatusCode))
	}

	return result, nil
}
