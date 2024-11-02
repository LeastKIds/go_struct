package mapper

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type DummyMapper struct{}

func NewDummyMapper() *DummyMapper {
	return &DummyMapper{}
}

func (m *DummyMapper) ToEntity(usecase *dto.Dummy) *entity.Dummy {
	return &entity.Dummy{
		ID:   usecase.ID,
		Name: usecase.Name,
		Age:  usecase.Age,
	}
}

func (m *DummyMapper) ToUsecase(entity *entity.Dummy) *dto.Dummy {
	return &dto.Dummy{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}
}
