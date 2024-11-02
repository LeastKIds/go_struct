package server

import (
	adpaterDummyHandler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	adapterDummyMapper "github.com/LeastKIds/go_struct/internal/adapter/mapper"
	infrastructureMapper "github.com/LeastKIds/go_struct/internal/infrastructure/database/mapper"
	infrastructureRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	infrastructureDummyRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository/dummy"
	usecaseDummyInteractor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
	usecaseDummyMapper "github.com/LeastKIds/go_struct/internal/usecase/mapper"
	usecaseMapper "github.com/LeastKIds/go_struct/internal/usecase/mapper"
)

func Server() {
	infrastructureDummyMapper := infrastructureMapper.NewInfrastructureDummyMapper()
	infrastructureDummyRepo := infrastructureDummyRepository.NewDummyRepository(infrastructureDummyMapper)
	infrastructureRepo := infrastructureRepository.NewInfrastructureRepository(infrastructureDummyRepo)

	usecaseDummyMapper := usecaseDummyMapper.NewUsecaseDummyMapper()
	usecaseMapper := usecaseMapper.NewUsecaseMapper(usecaseDummyMapper)
	usecaseDummyInteractor := usecaseDummyInteractor.NewInteractorDummy(*infrastructureRepo, *usecaseMapper)

	adpaterDummyMapper := adapterDummyMapper.NewAdapterDummyMapper()
	adpaterDummyHandler := adpaterDummyHandler.NewDummyHandler(usecaseDummyInteractor, adpaterDummyMapper)

}
