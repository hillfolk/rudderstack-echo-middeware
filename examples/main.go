package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rudderlabs/analytics-go/v4"
	rudder "github/hillfolk/rudderstack-echo-middleware"
	"net/http"
)

func main() {
	e := echo.New()
	client := rudder.NewClient("2KEB8SmnueBBHxM71ZiumKLMRZ7", "https://illuminarefzr.dataplane.rudderstack.com")
	e.Use(rudder.ContextRudderStack(client))
	e.GET("/", func(c echo.Context) error {
		client := c.Get(rudder.ContextRudderStackName).(rudder.IClient)
		client.Enqueue(analytics.Track{
			UserId: "1hKOmRA4GRlm",
			Event:  "Test Event",
		})
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
