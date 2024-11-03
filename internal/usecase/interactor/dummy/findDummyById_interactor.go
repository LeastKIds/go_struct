package dummy

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
	"gorm.io/gorm"
)

func (i *DummyInteractor) FindDummyById(db *gorm.DB, id int64) (*dto.Dummy, error) {
	dummy, err := i.repo.Dummy.FindBy(db, i.repo.Scopes.ByID(id))
	if err != nil {
		return nil, err
	}

	return mapper.DummyToDto(dummy), nil
}
