package server

import (
	dummyHandler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	cm "github.com/LeastKIds/go_struct/internal/adapter/middleware"
	"github.com/LeastKIds/go_struct/internal/adapter/router"
	dummyRouter "github.com/LeastKIds/go_struct/internal/adapter/router/dummy"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
	dummyRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository/dummy"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/scope"
	dummyInteractor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
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

	DI(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func DI(e *echo.Echo) {
	s := scope.NewScope()
	drepo := dummyRepository.NewDummyRepository()
	repo := repository.NewRepository(s, drepo)
	di := dummyInteractor.NewDummyInteractor(*repo)
	dh := dummyHandler.NewDummyHandler(di)
	dr := dummyRouter.NewDummyRouter(dh)
	r := router.NewRouter(dr)

	r.InitRouter(e)
}
