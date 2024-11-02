package server

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/labstack/echo/v4"
)

func Server() {
	config := config.NewConfig()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
