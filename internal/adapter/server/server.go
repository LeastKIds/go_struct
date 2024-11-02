package server

import (
	handler "github.com/LeastKIds/go_struct/internal/adapter/http/handler/dummy"
	"github.com/LeastKIds/go_struct/internal/adapter/http/router"
	dummyRouter "github.com/LeastKIds/go_struct/internal/adapter/http/router/dummy"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/repository"
	usecase "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() {
	config := config.NewConfig()

	db, err := connection.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}
	dummyRepo := repository.NewDummyRepository(db)
	dummyUsecase := usecase.NewDummyInteractor(dummyRepo)
	dummyHandler := handler.NewDummyHandler(dummyUsecase)
	dummyRouter := dummyRouter.NewDummyRouter(dummyHandler)
	router := router.NewRouter(dummyRouter)

	e := echo.New()

	e.Use(middleware.Logger())  // 모든 요청을 로깅
	e.Use(middleware.Recover()) // 패닉을 복구하여 에러를 로깅

	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
