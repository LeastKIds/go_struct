package dummy

import (
	interactor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	"github.com/labstack/echo/v4"
)

type IDummyHandler interface {
	GetDummy(c echo.Context) error
}

type DummyHandler struct {
	dummyUsecase interactor.IDummyInteractor
}

func NewDummyHandler(du interactor.IDummyInteractor) *DummyHandler {
	return &DummyHandler{dummyUsecase: du}
}
