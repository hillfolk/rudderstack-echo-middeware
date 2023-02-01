package rudderstack_echo_middeware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rudderlabs/analytics-go/v4"
)

type ContextRudderStackType string // Path: middleware.go

var ContextRudderStackName = "RUDDERSTACK" // Path: middleware.go

func ContextRudderSatck(key, urls string) echo.MiddlewareFunc {
	client := NewClient(key, urls)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			client := client.GetClient()
			c.SetRequest(req.WithContext(context.WithValue(ctx, ContextClientName, client)))
			return next(c)
		}
	}
}