package mysql

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mapper"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
	"gorm.io/gorm"
)

type Dummy struct {
	db *gorm.DB
}

func NewDummy(db *gorm.DB) *Dummy {
	return &Dummy{db: db}
}

func (d *Dummy) Create(entity *entity.Dummy) error {
	var model = mapper.NewDummyMapper().ToModel(entity)
	return d.db.Create(&model).Error
}

func (d *Dummy) FindAll(scopes ...repository.Scope) (*[]entity.Dummy, error) {
	var models *[]model.Dummy
	if err := d.db.Scopes(scopes...).Find(&models).Error; err != nil {
		return nil, err
	}
	return mapper.NewDummyMapper().ToEntities(models), nil
}

func (d *Dummy) FindBy(scopes ...repository.Scope) (*entity.Dummy, error) {
	var model model.Dummy
	if err := d.db.Scopes(scopes...).First(&model).Error; err != nil {
		return nil, err
	}
	return mapper.NewDummyMapper().ToEntity(&model), nil
}

func (d *Dummy) Update(entity *entity.Dummy) error {
	var model = mapper.NewDummyMapper().ToModel(entity)
	return d.db.Save(&model).Error
}

func (d *Dummy) Delete(scopes ...repository.Scope) error {
	return d.db.Scopes(scopes...).Delete(&model.Dummy{}).Error
}
