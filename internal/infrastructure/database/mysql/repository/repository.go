package repository

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/repository/dummy"
)

func NewDummyRepository() repository.ISubRepository[entity.Dummy] {
	return dummy.NewDummyRepository()
}
