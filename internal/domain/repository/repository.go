package repository

import (
	"gorm.io/gorm"
)

type Scope = func(db *gorm.DB) *gorm.DB

type Repo[model any] interface {
	FindBy(...Scope) (*model, error)
	FindAll(...Scope) (*[]model, error)
	Create(m *model) error
	Update(m *model) error
	Delete(...Scope) error
}
