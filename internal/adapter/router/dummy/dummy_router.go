package dummy

import (
	"github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	"github.com/LeastKIds/go_struct/internal/adapter/router"
	"github.com/labstack/echo/v4"
)

type DummyRouter struct {
	DummyHandler dummy.IDummyHandler
}

func NewDummyRouter(dummyHandler dummy.IDummyHandler) router.ISubRouter {
	return &DummyRouter{DummyHandler: dummyHandler}
}

func (r *DummyRouter) InitRouter(g *echo.Group) {
	g.GET("/:id", r.DummyHandler.GetDummy)
	g.POST("", r.DummyHandler.PostDummy)
}
