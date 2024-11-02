package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "user:password@tcp(localhost:3306)/go_struct_db?charset=utf8mb4&parseTime=True&loc=Local"

type Mysql struct{}

func NewMysql() *Mysql {
	return &Mysql{}
}

func (m *Mysql) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
