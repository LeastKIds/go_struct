package database

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql"
	"gorm.io/gorm"
)

type IConnection interface {
	Connect() (*gorm.DB, error)
}

func NewConnection(dsn string) IConnection {
	return mysql.NewMysql(dsn)
}
