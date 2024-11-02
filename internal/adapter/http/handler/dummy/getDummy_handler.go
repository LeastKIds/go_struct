package dummy

import (
	"strconv"

	"github.com/LeastKIds/go_struct/internal/adapter/http/mapper"
	interactor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	"github.com/labstack/echo/v4"
)

type IDummy interface {
	GetDummy(c echo.Context) error
}

type Dummy struct {
	dummyUsecase interactor.IDummy
}

func NewDummy(du interactor.IDummy) *Dummy {
	return &Dummy{dummyUsecase: du}
}

func (d *Dummy) GetDummy(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	entity, err := d.dummyUsecase.FindDummy(id)
	if err != nil {
		return err
	}

	res := mapper.NewDummyMapper().ToResponse(entity)

	return c.JSON(200, res)
}
