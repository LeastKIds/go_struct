package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	usecase "github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type IDummyInteractor interface {
	FindDummy(id int64) (*usecase.Dummy, error)
}

type DummyInteractor struct {
	DummyRepo repository.Repo[entity.Dummy]
}

func NewDummy(repo repository.Repo[entity.Dummy]) *DummyInteractor {
	return &DummyInteractor{DummyRepo: repo}
}
