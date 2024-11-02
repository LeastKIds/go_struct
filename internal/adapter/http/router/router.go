package router

import (
	dummyRouter "github.com/LeastKIds/go_struct/internal/adapter/http/router/dummy"
	"github.com/labstack/echo/v4"
)

type RouterStruct struct {
	dummyRouter dummyRouter.IDummyRouter
}

func NewRouter(dummyRouter dummyRouter.IDummyRouter) RouterStruct {
	return RouterStruct{
		dummyRouter: dummyRouter,
	}
}

func (rs *RouterStruct) Router(e *echo.Echo) {
	dummyGroup := e.Group("/dummy")
	{
		rs.dummyRouter.Router(dummyGroup)
	}
}
