package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/repository/mysql/scope"
)

type Dummy struct {
	DummyRepo repository.Repo[entity.Dummy]
}

func NewDummy(repo repository.Repo[entity.Dummy]) *Dummy {
	return &Dummy{DummyRepo: repo}
}

func (d *Dummy) FindDummy(id int64) (*entity.Dummy, error) {
	return d.DummyRepo.FindBy(scope.ByID(id))
}
