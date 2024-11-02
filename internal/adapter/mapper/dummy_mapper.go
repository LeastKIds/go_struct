package mapper

import (
	"github.com/LeastKIds/go_struct/internal/adapter/dto/response"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type IAdapterDummyMapper interface {
	ToResponse(dto *dto.Dummy) *response.Dummy
}

type AdapterDummyMapper struct{}

func NewAdapterDummyMapper() *AdapterDummyMapper {
	return &AdapterDummyMapper{}
}

func (m *AdapterDummyMapper) ToResponse(dto *dto.Dummy) *response.Dummy {
	return &response.Dummy{
		ID:   dto.ID,
		Name: dto.Name,
		Age:  dto.Age,
	}
}
