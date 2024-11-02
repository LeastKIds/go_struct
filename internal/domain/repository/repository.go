package repository

import "gorm.io/gorm"

type Scope = func(db *gorm.DB) *gorm.DB

type IRepository[T any] interface {
	FindBy(db *gorm.DB, scopes ...Scope) (*T, error)
	FindAll(db *gorm.DB, scopes ...Scope) (*[]T, error)
	Create(db *gorm.DB, model *T) error
	Update(db *gorm.DB, model *T) error
	Delete(db *gorm.DB, model *T) error
}
