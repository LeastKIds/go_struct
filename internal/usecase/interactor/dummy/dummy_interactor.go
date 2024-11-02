package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/scope"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
	"gorm.io/gorm"
)

type IInteractorDummy interface{}

type InteractorDummy struct {
	repo   repository.InfrastructureRepository
	mapper mapper.UsecaseMapper
}

func NewInteractorDummy(repo repository.InfrastructureRepository, mapper mapper.UsecaseMapper) *InteractorDummy {
	return &InteractorDummy{repo: repo, mapper: mapper}
}

func (i *InteractorDummy) GetDummy(db *gorm.DB, id int64) (*dto.Dummy, error) {
	dummy, err := i.repo.InfrastructureDummyRepo.FindBy(db, scope.ByID(id))
	if err != nil {
		return nil, err
	}

	return i.mapper.UsecaseDummyMapper.ToDTO(dummy), nil
}
