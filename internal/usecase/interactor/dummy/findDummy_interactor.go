package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/repository/mysql/scope"
	usecase "github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
)

type IDummy interface {
	FindDummy(id int64) (*usecase.Dummy, error)
}

type Dummy struct {
	DummyRepo repository.Repo[entity.Dummy]
}

func NewDummy(repo repository.Repo[entity.Dummy]) *Dummy {
	return &Dummy{DummyRepo: repo}
}

func (d *Dummy) FindDummy(id int64) (*usecase.Dummy, error) {
	dummy, err := d.DummyRepo.FindBy(scope.ByID(id))
	if err != nil {
		return nil, err
	}

	return mapper.NewDummyMapper().ToUsecase(dummy), nil
}
