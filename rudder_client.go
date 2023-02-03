package rudderstack_echo_middeware

import (
	"github.com/rudderlabs/analytics-go/v4"
)

type IClient interface {
	GetClient() analytics.Client
	Enqueue(msg analytics.Message) error
	Close() error
}
type Client struct {
	rudderClient analytics.Client
}

func NewClient(key, dataPanelUrl string) IClient {
	client := analytics.New(key, dataPanelUrl)
	return &Client{
		rudderClient: client,
	}
}

func (c *Client) GetClient() analytics.Client {
	return c.rudderClient
}

func (c *Client) Enqueue(msg analytics.Message) error {
	return c.rudderClient.Enqueue(msg)
}

func (c *Client) Close() error {
	return c.rudderClient.Close()
}
