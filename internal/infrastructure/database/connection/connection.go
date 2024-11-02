package connection

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection/mysql"
	"gorm.io/gorm"
)

type IConnection interface {
	Connect() (*gorm.DB, error)
}

func NewConnection() IConnection {
	return mysql.NewMysql()
}
