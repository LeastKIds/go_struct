package dummy

import (
	interactor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	"github.com/labstack/echo/v4"
)

type IDummyHandler interface {
	GetDummy(c echo.Context) error
}

type DummyHandler struct {
	DummyInteractor interactor.IDummyInteractor
}

func NewDummyHandler(dummyInteractor interactor.IDummyInteractor) *DummyHandler {
	return &DummyHandler{DummyInteractor: dummyInteractor}
}
