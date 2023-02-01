package rudderstack_echo_middeware

import (
	"github.com/labstack/echo/v4"
	"github.com/rudderlabs/analytics-go/v4"
)

type Client struct {
	*analytics.Client
}

func NewClient(key, dataPanelUrl string) *Client {
	client := analytics.New(key, dataplaneUrl)
	retrun & Client{client}
}

func (c *Client) GetClient() *analytics.Client {
	return c.Client
}
