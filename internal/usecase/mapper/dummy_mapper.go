package mapper

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type IUsecaseDummyMapper interface {
	ToDTO(entity *entity.Dummy) *dto.Dummy
}

type UsecaseDummyMapper struct{}

func NewUsecaseDummyMapper() *UsecaseDummyMapper {
	return &UsecaseDummyMapper{}
}

func (m *UsecaseDummyMapper) ToDTO(entity *entity.Dummy) *dto.Dummy {
	return &dto.Dummy{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}
}
