package server

import (
	adpaterDummyHandler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	adapterDummyMapper "github.com/LeastKIds/go_struct/internal/adapter/mapper"
	"github.com/LeastKIds/go_struct/internal/adapter/router"
	adapterDummyRaouter "github.com/LeastKIds/go_struct/internal/adapter/router/dummy"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
	infrastructureMapper "github.com/LeastKIds/go_struct/internal/infrastructure/database/mapper"
	infrastructureRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	infrastructureDummyRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository/dummy"
	usecaseDummyInteractor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	usecaseMapper "github.com/LeastKIds/go_struct/internal/usecase/mapper"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Start() {
	e := echo.New()

	infrastructureDummyMapper := infrastructureMapper.NewInfrastructureDummyMapper()
	infrastructureDummyRepo := infrastructureDummyRepository.NewDummyRepository(infrastructureDummyMapper)
	infrastructureRepo := infrastructureRepository.NewInfrastructureRepository(infrastructureDummyRepo)

	usecaseDummyMapper := usecaseMapper.NewUsecaseDummyMapper()
	usecaseMapper := usecaseMapper.NewUsecaseMapper(usecaseDummyMapper)
	usecaseDummyInteractor := usecaseDummyInteractor.NewInteractorDummy(*infrastructureRepo, *usecaseMapper)

	adpaterDummyMapper := adapterDummyMapper.NewAdapterDummyMapper()
	adpaterDummyHandler := adpaterDummyHandler.NewDummyHandler(usecaseDummyInteractor, adpaterDummyMapper)
	adapterDummyRouter := adapterDummyRaouter.NewDummyRouter(adpaterDummyHandler)
	router := router.NewRouter(e, adapterDummyRouter)

	config := config.NewConfig()
	db, err := database.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		e.Logger.Fatal(err)
	}
	server := NewServer(router, db)
	server.InitServer()

	e.Logger.Fatal(e.Start(":8080"))
}

type IServer interface {
	InitServer()
}

type Server struct {
	router router.IRouter
	db     *gorm.DB
}

func NewServer(router router.IRouter, db *gorm.DB) *Server {
	return &Server{router: router, db: db}
}

func (s *Server) InitServer() {
	s.router.InitRouter()
}
