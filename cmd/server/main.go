package main

import (
	"github.com/WanderningMaster/rpi-dash/components"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "./static")

	e.GET("/", func(c echo.Context) error {
		return components.Show().Render(c.Request().Context(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
