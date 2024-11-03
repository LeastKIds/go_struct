package framework

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type EchoFramework struct {
	echo *echo.Echo
	db   *gorm.DB
}

func NewEchoFramework(db *gorm.DB) IFramework {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &EchoFramework{echo: e, db: db}
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
	return &EchoGroup{group: grp, db: e.db}
}

type EchoGroup struct {
	group *echo.Group
	db    *gorm.DB
}

func (g *EchoGroup) GET(path string, handler HandlerFuncWithDB) {
	g.group.GET(path, func(c echo.Context) error {
		transaction(g.db, c, handler)
		return nil
	})
}

func (g *EchoGroup) POST(path string, handler HandlerFuncWithDB) {
	g.group.POST(path, func(c echo.Context) error {
		transaction(g.db, c, handler)
		return nil
	})
}

func (g *EchoGroup) PUT(path string, handler HandlerFuncWithDB) {
	g.group.PUT(path, func(c echo.Context) error {
		transaction(g.db, c, handler)
		return nil
	})
}

func (g *EchoGroup) DELETE(path string, handler HandlerFuncWithDB) {
	g.group.DELETE(path, func(c echo.Context) error {
		transaction(g.db, c, handler)
		return nil
	})
}

func (g *EchoGroup) Group(prefix string) IGroup {
	grp := g.group.Group(prefix)
	return &EchoGroup{group: grp, db: g.db}
}

func transaction(db *gorm.DB, c echo.Context, handler HandlerFuncWithDB) {
	if c.Request().Method == http.MethodGet {
		handler(c.Response().Writer, c.Request(), db)
		return
	}

	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, tx.Error)
		return
	}

	handler(c.Response().Writer, c.Request(), tx)

	if tx.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, tx.Error)
		return
	}

	tx.Commit()
}
