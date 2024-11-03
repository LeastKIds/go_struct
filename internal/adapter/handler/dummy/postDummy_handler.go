package dummy

import (
	"github.com/LeastKIds/go_struct/internal/adapter/dto/request"
	"github.com/LeastKIds/go_struct/internal/adapter/mapper"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *DummyHandler) PostDummy(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	req := request.NewDummy()
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err)
	}

	dummyDto := mapper.DummyToDto(req)

	if err := h.DummyInteractor.CreateDummy(db, dummyDto); err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, "success")
}
