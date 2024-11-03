package dummy

import "github.com/labstack/echo/v4"

type IDummyHandler interface {
	GetDummy(c echo.Context) error
}

type DummyHandler struct {
	DummyUsecase usecase.IDummyUsecase
}

func NewDummyHandler(dummyUsecase *usecase.IDummyUsecase) *DummyHandler {
	return &DummyHandler{DummyUsecase: dummyUsecase}
}
