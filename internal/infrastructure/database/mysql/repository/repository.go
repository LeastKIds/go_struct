package repository

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
)

type InfrastructureRepository struct {
	InfrastructureDummyRepo repository.IRepository[entity.Dummy]
}

func NewInfrastructureRepository(infrastructureDummyRepo repository.IRepository[entity.Dummy]) *InfrastructureRepository {
	return &InfrastructureRepository{InfrastructureDummyRepo: infrastructureDummyRepo}
}
