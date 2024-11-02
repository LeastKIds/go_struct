package router

import (
	dummyRouter "github.com/LeastKIds/go_struct/internal/adapter/http/router/dummy"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	dummyGroup := e.Group("/dummy")
	{
		dummyRouter.Router(dummyGroup)
	}
}
