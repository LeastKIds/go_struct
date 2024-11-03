package dummy

import (
	"strconv"

	"github.com/LeastKIds/go_struct/internal/adapter/mapper"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *DummyHandler) GetDummy(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(400, "Invalid ID")
	}

	dummy, err := h.DummyInteractor.FindDummyById(db, id)
	if err != nil {
		return c.JSON(500, "Internal Server Error")
	}

	res := mapper.DummyToResponse(dummy)

	return c.JSON(200, res)
}
