package server

import (
	cm "github.com/LeastKIds/go_struct/internal/adapter/middleware"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	cfg := config.NewConfig()
	db, err := database.NewConnection(cfg.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(cm.TransactionMiddleware(db))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := DI(e); err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}

func DI(e *echo.Echo) error {
	r, err := InitializeRouter()
	if err != nil {
		return err
	}

	r.InitRouter(e)
	return nil
}
