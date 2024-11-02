package mapper

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

type IInfrastructureDummyMapper interface {
	ToEntity(model *model.Dummy) *entity.Dummy
	ToModel(entity *entity.Dummy) *model.Dummy
}

type InfrastructureDummyMapper struct{}

func NewInfrastructureDummyMapper() *InfrastructureDummyMapper {
	return &InfrastructureDummyMapper{}
}

func (m *InfrastructureDummyMapper) ToEntity(model *model.Dummy) *entity.Dummy {
	return &entity.Dummy{
		ID:   model.ID,
		Name: model.Name,
		Age:  model.Age,
	}
}

func (m *InfrastructureDummyMapper) ToModel(entity *entity.Dummy) *model.Dummy {
	return &model.Dummy{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}
}
