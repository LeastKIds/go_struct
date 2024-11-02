package dummy

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *DummyHandler) GetDummy(c echo.Context) error {
	var db *gorm.DB

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	dummy, err := h.InteractorDummy.GetDummy(db, id)
	if err != nil {
		return err
	}

	return c.JSON(200, h.mapper.ToResponse(dummy))
}
