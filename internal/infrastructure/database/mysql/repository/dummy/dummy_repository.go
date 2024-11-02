package dummy

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mapper"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type DummyRepository struct {
	infrastructureDummyMapper mapper.IInfrastructureDummyMapper
}

func NewDummyRepository(infrastructureMapper mapper.IInfrastructureDummyMapper) *DummyRepository {
	return &DummyRepository{infrastructureDummyMapper: infrastructureMapper}
}

func (r *DummyRepository) FindBy(db *gorm.DB, scopes ...repository.Scope) (*entity.Dummy, error) {
	var model model.Dummy
	if err := db.Scopes(scopes...).First(&model).Error; err != nil {
		return nil, err
	}

	return r.infrastructureDummyMapper.ToEntity(&model), nil
}

func (r *DummyRepository) Create(db *gorm.DB, entity *entity.Dummy) error {
	model := r.infrastructureDummyMapper.ToModel(entity)
	if err := db.Create(model).Error; err != nil {
		return err
	}

	return nil
}
