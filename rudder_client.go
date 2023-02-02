package rudderstack_echo_middeware

import (
	"github.com/rudderlabs/analytics-go/v4"
)

type Client struct {
	RudderClient analytics.Client
}

func NewClient(key, dataPanelUrl string) *Client {
	client := analytics.New(key, dataPanelUrl)
	return &Client{
		RudderClient: client,
	}
}

func (c *Client) GetClient() analytics.Client {
	return c.RudderClient
}

func (c *Client) Enqueue(msg analytics.Message) error {
	return c.RudderClient.Enqueue(msg)
}

func (c *Client) Close() error {
	return c.RudderClient.Close()
}
