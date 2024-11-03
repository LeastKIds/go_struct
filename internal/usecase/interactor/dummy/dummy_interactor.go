package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

type IDummyInteractor interface {
	FindDummyById(db *gorm.DB, id int64) (*dto.Dummy, error)
	CreateDummy(db *gorm.DB, dummy *dto.Dummy) error
}

type DummyInteractor struct {
	repo repository.Repository
}

func NewDummyInteractor(repo repository.Repository) *DummyInteractor {
	return &DummyInteractor{repo: repo}
}
