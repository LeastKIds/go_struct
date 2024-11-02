package database

import "gorm.io/gorm"

type IConnection interface {
	Connect(dsn string) (*gorm.DB, error)
}
