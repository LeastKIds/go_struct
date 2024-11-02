package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/scope"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

type IInteractorDummy interface{}

type InteractorDummy struct {
	repo repository.InfrastructureRepository
}

func NewInteractorDummy(repo repository.InfrastructureRepository) *InteractorDummy {
	return &InteractorDummy{repo: repo}
}

func (i *InteractorDummy) GetDummy(db *gorm.DB, id int64) (*dto.Dummy, error) {
	dummy, err := i.repo.InfrastructureDummyRepo.FindBy(db, scope.ByID(id))
	if err != nil {
		return nil, err
	}
}
