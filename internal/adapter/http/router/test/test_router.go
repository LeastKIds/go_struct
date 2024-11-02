package test

import (
	handler "github.com/LeastKIds/go_struct/internal/adapter/http/handler/test"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Group) {
	e.GET("", handler.Test)
}
