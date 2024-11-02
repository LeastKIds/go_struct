package server

import "github.com/labstack/echo/v4"

func Server() {
	start()
}

func start() {
	e := echo.New()

	v1 := e.Group("/v1")

	auth := v1.Group("/auth")
	{
		authRoutes.Routes(auth)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
