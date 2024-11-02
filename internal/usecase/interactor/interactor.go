package interactor

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/usecase/interactor/dummy"
)

type IInteractor interface {
	DummyInteractor(repo repository.Repo[entity.Dummy]) dummy.IDummyInteractor
}

type Interactor struct {
	repo            repository.Repo[entity.Dummy]
	dummyInteractor dummy.IDummyInteractor
}

func NewInteractor(repo repository.Repo[entity.Dummy]) *Interactor {
	return &Interactor{repo: repo}
}

func (i *Interactor) DummyInteractor() dummy.IDummyInteractor {
	if i.dummyInteractor == nil {
		i.dummyInteractor = dummy.NewDummyInteractor(i.repo)
	}
	return i.dummyInteractor
}
