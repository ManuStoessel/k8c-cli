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
	"net/http"
	"net/url"
)

const (
	apiV1Path string = "/api/v1"
)

// Client holds all config and the http.Client needed to talk to the Kubermatic API
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	token      string
}

// NewClient creates a new Client for the Kubermatic API
func NewClient(baseurl string, token string) (*Client, error) {
	parsedurl, err := url.Parse(baseurl + apiV1Path)
	if err != nil {
		return &Client{}, err
	}

	httpClient := &http.Client{}

	client := &Client{}
	client.BaseURL = parsedurl
	client.httpClient = httpClient
	client.token = token

	return client, nil
}
