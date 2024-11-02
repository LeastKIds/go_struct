package repository

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/repository/mysql"
	"gorm.io/gorm"
)

func NewDummyRepository(db *gorm.DB) repository.Repo[entity.Dummy] {
	return mysql.NewDummy(db)
}
