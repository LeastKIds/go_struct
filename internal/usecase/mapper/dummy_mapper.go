package mapper

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

func DummyToDto(entity *entity.Dummy) *dto.Dummy {
	return &dto.Dummy{
		ID:   entity.ID,
		Name: entity.Name,
		Age:  entity.Age,
	}
}

func DummyToEntity(dto *dto.Dummy) *entity.Dummy {
	return &entity.Dummy{
		ID:   dto.ID,
		Name: dto.Name,
		Age:  dto.Age,
	}
}
