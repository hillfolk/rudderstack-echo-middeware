package rudderstack_echo_middeware

import (
	"context"
	"github.com/labstack/echo/v4"
)

var ContextRudderStackName = "RUDDERSTACK" // Path: middleware.go

func ContextRudderStack(client *Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()
			c.SetRequest(req.WithContext(context.WithValue(ctx, ContextRudderStackName, client)))
			return next(c)
		}
	}
}
