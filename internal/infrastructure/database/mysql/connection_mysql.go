package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct{ dsn string }

func NewMysql(dsn string) *Mysql {
	return &Mysql{dsn: dsn}
}

func (m *Mysql) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
