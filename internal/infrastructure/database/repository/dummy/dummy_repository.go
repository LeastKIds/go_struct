package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mapper"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type DummyRepository struct{}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{}
}

func (r *DummyRepository) FindBy(db *gorm.DB, scopes ...repository.Scope) (*entity.Dummy, error) {
	var model model.Dummy
	if err := db.Scopes(scopes...).First(&model).Error; err != nil {
		return nil, err
	}

	return mapper.NewInfrastructureDummyMapper()
}
