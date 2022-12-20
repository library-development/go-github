package github

import "net/http"

type Client struct {
}

func request(method, path, body any) (http.Response, error) {
}

func (c *Client) CreateRepo(org, name string, public bool) error {
}
