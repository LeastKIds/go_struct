package dummy

import (
	"github.com/LeastKIds/go_struct/internal/adapter/mapper"
	"github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	"github.com/labstack/echo/v4"
)

type IDummyHandler interface {
	GetDummy(c echo.Context) error
}

type DummyHandler struct {
	InteractorDummy dummy.IInteractorDummy
	mapper          mapper.IAdapterDummyMapper
}

func NewDummyHandler(interactorDummy dummy.IInteractorDummy, mapper mapper.IAdapterDummyMapper) *DummyHandler {
	return &DummyHandler{InteractorDummy: interactorDummy, mapper: mapper}
}
