package dummy

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *DummyHandler) GetDummy(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	return c.JSON(200, "Hello World")
}
