package server

import "github.com/labstack/echo/v4"

func Server() {
	start()
}

func start() {
	e := echo.New()

	v1 := e.Group("/v1")

	test := v1.Group("/test")
	{
		testRoutes.Router(test)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
