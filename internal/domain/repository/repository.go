package repository

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

type IScopes interface {
	ByID(id int64) Scope
}

type Scope = func(db *gorm.DB) *gorm.DB

type ISubRepository[T any] interface {
	FindBy(db *gorm.DB, scopes ...Scope) (*T, error)
}

type Repository struct {
	Scopes IScopes
	Dummy  ISubRepository[dto.Dummy]
}

func NewRepository(scopes IScopes, dummy ISubRepository[dto.Dummy]) *Repository {
	return &Repository{Scopes: scopes, Dummy: dummy}
}
