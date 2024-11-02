package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

func main() {
	config := config.NewConfig()
	db, err := database.NewConnection(config.DATABASE_URL).Connect()
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
