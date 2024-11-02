package connection

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection/mysql"
	"gorm.io/gorm"
)

type IConnection interface {
	Connect() (*gorm.DB, error)
}

type Connection struct{ dsn string }

func NewConnection(dsn string) IConnection {
	return mysql.NewMysql(dsn)
}
