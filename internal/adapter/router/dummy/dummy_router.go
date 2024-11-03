package dummy

import (
	handler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	"github.com/labstack/echo/v4"
)

type DummyRouter struct {
	handler handler.IDummyHandler
}

func NewDummyRouter(handler handler.IDummyHandler) *DummyRouter {
	return &DummyRouter{handler: handler}
}

func (r *DummyRouter) Routes(g *echo.Group) {
	g.GET("/:id", r.handler.GetDummy)
}
