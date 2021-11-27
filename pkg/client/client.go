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
