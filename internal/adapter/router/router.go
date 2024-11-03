package router

import "github.com/labstack/echo/v4"

type IRouter interface {
	InitRouter(e *echo.Echo)
}

type ISubRouter interface {
	InitRouter(g *echo.Group)
}

type Router struct {
	DummyRouter ISubRouter
}

func NewRouter(dummyRouter ISubRouter) IRouter {
	return &Router{DummyRouter: dummyRouter}
}

func (r *Router) InitRouter(e *echo.Echo) {
	v1 := e.Group("/v1")

	dummyGroup := v1.Group("/dummy")
	{
		r.DummyRouter.InitRouter(dummyGroup)
	}

}
