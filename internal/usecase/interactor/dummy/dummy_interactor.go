package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type IInteractorDummy interface{}

type InteractorDummy struct{}

func NewInteractorDummy() *InteractorDummy {
	return &InteractorDummy{}
}

func (i *InteractorDummy) GetDummy(repo repository.InfrastructureRepository, id int64) (*dto.Dummy, error) {
	dummy, err := repo.InfrastructureDummyRepo.FindBy(scope.ById(id))
	if err != nil {
		return nil, err
	}
}
