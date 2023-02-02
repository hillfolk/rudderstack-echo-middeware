package main

import (
	"net/http"
	"github.com/hillfolk/rudderstack-echo-middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	client := rudderstack_echo_middeware.NewClient("key", "url")
	e.Use(client)
	e.GET("/", func(c echo.Context) error {
		client := c.Get(ContextRudderStackName)
		client.Enqueue(rudderstack.Track{
			client.Enqueue(analytics.Track{
				UserId: "1hKOmRA4GRlm",
				Event:  "Test Event",
			}),
		})

	}
	return c.String(http.StatusOK, "Hello, World!")
})
e.Logger.Fatal(e.Start(":1323"))
}
