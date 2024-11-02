package test

import (
	"fmt"
	"net/http"

	"github.com/LeastKIds/go_struct/internal/usecase/interactor/test"
	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	fmt.Println("handler")
	testInteractor := test.GetTestData()
	return c.String(http.StatusOK, testInteractor)
}
