package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/kubermatic/go-kubermatic/models"
)

const (
	apiV1Path   string = "/api/v1"
	projectPath string = "/projects"
)

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	token      string
}

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

func (c *Client) ListProjects() ([]*models.Project, error) {
	return []*models.Project{}, nil
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
