package mapper

import (
	"github.com/LeastKIds/go_struct/internal/adapter/dto/response"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type IUsecaseDummyMapper interface {
	ToResponse(dto *dto.Dummy) *response.Dummy
}

type UsecaseDummyMapper struct{}

func NewUsecaseDummyMapper() *UsecaseDummyMapper {
	return &UsecaseDummyMapper{}
}

func (m *UsecaseDummyMapper) ToResponse(dto *dto.Dummy) *response.Dummy {
	return &response.Dummy{
		ID:   dto.ID,
		Name: dto.Name,
		Age:  dto.Age,
	}
}
