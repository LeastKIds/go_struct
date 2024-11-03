package framework

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoFramework struct {
	echo *echo.Echo
}

func NewEchoFramework() IFramework {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &EchoFramework{echo: e}
}

func (e *EchoFramework) Use(middleware func(http.Handler) http.Handler) {
	e.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c.SetRequest(r.WithContext(c.Request().Context()))
				next(c)
			})).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}
	})
}

func (e *EchoFramework) Start(address string) error {
	return e.echo.Start(address)
}

func (e *EchoFramework) Group(prefix string) IGroup {
	grp := e.echo.Group(prefix)
	return &EchoGroup{group: grp}
}

type EchoGroup struct {
	group *echo.Group
}

func (g *EchoGroup) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.GET(path, func(c echo.Context) error {
		handler(c.Response().Writer, c.Request())
		return nil
	})
}

func (g *EchoGroup) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.POST(path, func(c echo.Context) error {
		handler(c.Response().Writer, c.Request())
		return nil
	})
}

func (g *EchoGroup) PUT(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.PUT(path, func(c echo.Context) error {
		handler(c.Response().Writer, c.Request())
		return nil
	})
}

func (g *EchoGroup) DELETE(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.DELETE(path, func(c echo.Context) error {
		handler(c.Response().Writer, c.Request())
		return nil
	})
}

func (g *EchoGroup) Group(prefix string) IGroup {
	grp := g.group.Group(prefix)
	return &EchoGroup{group: grp}
}
