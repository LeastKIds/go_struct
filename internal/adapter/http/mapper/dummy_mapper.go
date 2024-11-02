package mapper

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

type DummyMapper struct{}

func NewDummyMapper() *DummyMapper {
	return &DummyMapper{}
}

func (m *DummyMapper) ToResponse(usecase *dto.Dummy) *dto.Dummy {
	return &dto.Dummy{
		ID:   usecase.ID,
		Name: usecase.Name,
		Age:  usecase.Age,
	}
}
