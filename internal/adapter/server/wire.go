//go:build wireinject
// +build wireinject

package server

import (
	"github.com/LeastKIds/go_struct/internal/adapter/router"
	"github.com/google/wire"
)

func InitializeRouter() (router.IRouter, error) {
	wire.Build(
		ProviderScopes,
		ProviderDummyRepository,
		ProviderRepository,
		ProviderDummyInteractor,
		ProviderDummyHandler,
		ProviderDummyRouter,
		ProviderRouter,
	)
	return &router.Router{}, nil
}
