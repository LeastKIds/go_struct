package dummy

import (
	handler "github.com/LeastKIds/go_struct/internal/adapter/http/handler/dummy"
	"github.com/labstack/echo/v4"
)

type IDummyRouter interface {
	Router(g *echo.Group)
}

type DummyRouter struct {
	dummyHandler handler.IDummyHandler
}

func NewDummyRouter(dh handler.IDummyHandler) *DummyRouter {
	return &DummyRouter{dummyHandler: dh}
}

func (dr *DummyRouter) Router(g *echo.Group) {
	g.GET("", dr.dummyHandler.GetDummy)
}
