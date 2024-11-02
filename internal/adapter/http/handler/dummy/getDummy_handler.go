package dummy

import (
	"strconv"

	"github.com/LeastKIds/go_struct/internal/adapter/http/mapper"
	"github.com/labstack/echo/v4"
)

func (dh *DummyHandler) GetDummy(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	entity, err := dh.dummyUsecase.FindDummy(id)
	if err != nil {
		return err
	}

	res := mapper.NewDummyMapper().ToResponse(entity)

	return c.JSON(200, res)
}
