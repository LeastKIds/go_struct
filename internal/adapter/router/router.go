package router

import (
	"github.com/labstack/echo/v4"
)

type IRouter interface {
	InitRouter()
}

type ISubRouter interface {
	Routes(g *echo.Group)
}

type Router struct {
	e           *echo.Echo
	dummyRouter ISubRouter
}

func NewRouter(e *echo.Echo, dummyRouter ISubRouter) *Router {
	return &Router{e: e, dummyRouter: dummyRouter}
}

func (r *Router) InitRouter() {
	v1 := r.e.Group("/v1")

	dummyGroup := v1.Group("/dummy")
	{
		r.dummyRouter.Routes(dummyGroup)
	}
}
