package main

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
)

var Models = []interface{}{
	&entity.Dummy{},
}

func main() {
	config := config.NewConfig()
	db, err := database.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(Models...); err != nil {
		panic(err)
	}
}
