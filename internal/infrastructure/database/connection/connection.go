package connection

import "gorm.io/gorm"

type IConnection interface {
	Connect() (*gorm.DB, error)
}
