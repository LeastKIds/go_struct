package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

func main() {
	config := config.NewConfig()
	db, err := connection.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	dummy := &model.Dummy{
		Name: "Dummy",
		Age:  20,
	}

	if err := db.Create(dummy).Error; err != nil {
		panic(err)
	}
}
