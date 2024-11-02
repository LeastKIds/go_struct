package mapper

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

type DummyMapper struct{}

func NewDummyMapper() *DummyMapper {
	return &DummyMapper{}
}

func (m *DummyMapper) ToModel(entity *entity.Dummy) *model.Dummy {
	return &model.Dummy{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}
}

func (m *DummyMapper) ToEntity(model *model.Dummy) *entity.Dummy {
	return &entity.Dummy{
		ID:   model.ID,
		Name: model.Name,
		Age:  model.Age,
	}
}

func (m *DummyMapper) ToEntities(models *[]model.Dummy) *[]entity.Dummy {
	var entities []entity.Dummy
	for _, model := range *models {
		entities = append(entities, *m.ToEntity(&model))
	}
	return &entities
}
