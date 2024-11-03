package dummy

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
	"gorm.io/gorm"
)

func (i *DummyInteractor) CreateDummy(db *gorm.DB, dummyDto *dto.Dummy) error {
	dummyEntity := mapper.DummyToEntity(dummyDto)
	if err := i.repo.Dummy.Create(db, dummyEntity); err != nil {
		return err
	}

	return nil
}
