package scope

import (
	"github.com/LeastKIds/go_struct/internal/domain/repository"
	"gorm.io/gorm"
)

func (s *Scope) ByID(id int64) repository.Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
