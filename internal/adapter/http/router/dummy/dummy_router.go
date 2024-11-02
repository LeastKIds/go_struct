package dummy

import "github.com/labstack/echo/v4"

func Router(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(200, "Dummy")
	})
}
