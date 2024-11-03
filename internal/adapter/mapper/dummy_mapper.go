package mapper

import (
	"github.com/LeastKIds/go_struct/internal/adapter/dto/request"
	"github.com/LeastKIds/go_struct/internal/adapter/dto/response"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
)

func DummyToResponse(dto *dto.Dummy) *response.Dummy {
	return &response.Dummy{
		ID:   dto.ID,
		Name: dto.Name,
		Age:  dto.Age,
	}
}

func DummyToDto(request *request.Dummy) *dto.Dummy {
	return &dto.Dummy{
		Name: request.Name,
		Age:  request.Age,
	}
}
