package repository

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"gorm.io/gorm"
)

type IScopes interface {
	ByID(id int64) Scope
}

type Scope = func(db *gorm.DB) *gorm.DB

type ISubRepository[T any] interface {
	FindBy(db *gorm.DB, scopes ...Scope) (*T, error)
	Create(db *gorm.DB, entity *T) error
}

type Repository struct {
	Scopes IScopes
	Dummy  ISubRepository[entity.Dummy]
}

func NewRepository(scopes IScopes, dummy ISubRepository[entity.Dummy]) *Repository {
	return &Repository{Scopes: scopes, Dummy: dummy}
}
