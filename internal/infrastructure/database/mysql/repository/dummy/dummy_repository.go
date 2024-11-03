package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"gorm.io/gorm"
)

type DummyRepository struct {
}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{}
}

func (r *DummyRepository) FindBy(db *gorm.DB, scopes ...repository.Scope) (*entity.Dummy, error) {
	var dummy entity.Dummy

	if err := db.Scopes(scopes...).First(&dummy).Error; err != nil {
		return nil, err
	}

	return &dummy, nil
}

func (r *DummyRepository) Create(db *gorm.DB, entity *entity.Dummy) error {
	return db.Create(&entity).Error
}
