package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	apiV1Path   string = "/api/v1"
	projectPath string = apiV1Path + "/projects"
)

// Client holds all config and the http.Client needed to talk to the Kubermatic API
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	token      string
}

// NewClient creates a new Client for the Kubermatic API
func NewClient(baseurl string, token string) (*Client, error) {
	parsedurl, err := url.Parse(baseurl)
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

// ListProjects lists all projects a user has permission to see
func (c *Client) ListProjects() ([]*models.Project, error) {
	req, err := c.newRequest("GET", projectPath, nil)
	if err != nil {
		return nil, err
	}

	var result []*models.Project

	resp, err := c.do(req, result)
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

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	bearer := "Bearer " + c.token
	req.Header.Add("Authorization", bearer)
	return req, nil
}
