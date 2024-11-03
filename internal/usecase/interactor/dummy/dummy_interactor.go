package dummy

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

type IDummyInteractor interface {
	FindDummyById(db *gorm.DB, id int64) (*dto.Dummy, error)
}

type DummyInteractor struct {
	repo repository.IRepository
}

func NewDummyInteractor(repo repository.IRepository) *DummyInteractor {
	return &DummyInteractor{repo: repo}
}
