package test

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	fmt.Println("handler")
	return c.String(http.StatusOK, "Test")
}
