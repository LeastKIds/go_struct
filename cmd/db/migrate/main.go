package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

var Models = []interface{}{
	&model.Dummy{},
}

func main() {
	config := config.NewConfig()
	db, err := connection.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(Models...); err != nil {
		panic(err)
	}
}
