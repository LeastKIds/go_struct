package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
	"gorm.io/gorm"
)

type IInteractorDummy interface {
	GetDummy(db *gorm.DB, id int64) (*dto.Dummy, error)
}

type InteractorDummy struct {
	repo   repository.InfrastructureRepository
	mapper mapper.UsecaseMapper
}

func NewInteractorDummy(repo repository.InfrastructureRepository, mapper mapper.UsecaseMapper) *InteractorDummy {
	return &InteractorDummy{repo: repo, mapper: mapper}
}
