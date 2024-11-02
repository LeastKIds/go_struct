package dummy

import (
	handler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	"github.com/labstack/echo/v4"
)

type DummyRouter struct {
	g       *echo.Group
	handler handler.IDummyHandler
}

func NewDummyRouter(g *echo.Group, handler handler.IDummyHandler) *DummyRouter {
	return &DummyRouter{g: g, handler: handler}
}

func (r *DummyRouter) Routes() {
	r.g.GET("/:id", r.handler.GetDummy)
}
