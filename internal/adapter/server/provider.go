package server

import (
	dummyHandler "github.com/LeastKIds/go_struct/internal/adapter/handler/dummy"
	"github.com/LeastKIds/go_struct/internal/adapter/router"
	dummyRouter "github.com/LeastKIds/go_struct/internal/adapter/router/dummy"
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	dummyRepository "github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository/dummy"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/scope"
	dummyInteractor "github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
)

func ProviderScopes() repository.IScopes {
	return scope.NewScope()
}

func ProviderDummyRepository() repository.ISubRepository[entity.Dummy] {
	return dummyRepository.NewDummyRepository()
}

func ProviderRepository(scopes repository.IScopes, dummy repository.ISubRepository[entity.Dummy]) *repository.Repository {
	return repository.NewRepository(scopes, dummy)
}

func ProviderDummyInteractor(repo *repository.Repository) dummyInteractor.IDummyInteractor {
	return dummyInteractor.NewDummyInteractor(*repo)
}

func ProviderDummyHandler(di dummyInteractor.IDummyInteractor) dummyHandler.IDummyHandler {
	return dummyHandler.NewDummyHandler(di)
}

func ProviderDummyRouter(dh dummyHandler.IDummyHandler) router.ISubRouter {
	return dummyRouter.NewDummyRouter(dh)
}

func ProviderRouter(dr router.ISubRouter) router.IRouter {
	return router.NewRouter(dr)
}
