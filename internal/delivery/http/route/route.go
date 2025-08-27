package route

import (
	"iot-sensor/internal/delivery/http"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	App              *echo.Echo
	SensorController *http.SensorController
}

func (c *RouteConfig) Setup() {
	c.App.PATCH("/api/change", c.SensorController.ChangeInterval)
}
