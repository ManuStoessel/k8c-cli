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
)

const (
	datacenterPath string = ".." + apiV1Path + "/dc"
)

// ListDatacenters lists all datacenters that are available
func (c *Client) ListDatacenters() ([]models.Datacenter, error) {
	req, err := c.newRequest("GET", datacenterPath, nil)
	if err != nil {
		return nil, err
	}

	result := make([]models.Datacenter, 0)

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
