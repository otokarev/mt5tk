package client

import "github.com/otokarev/mt5tk/pkg/connection"

type Client struct {
	Connection *connection.Connection
}

func (c *Client) Get(url string) ([]byte, error) {
	return c.Connection.Get(url)
}

func (c *Client) Ping() {
	c.Connection.Ping()
}

func (c *Client) Post(url string, body []byte) ([]byte, error) {
	return c.Connection.Post(url, body)
}
