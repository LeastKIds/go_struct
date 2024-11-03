package main

import (
	"github.com/LeastKIds/go_struct/internal/domain/entity"
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
)

func main() {
	config := config.NewConfig()
	db, err := database.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	dummy := &entity.Dummy{
		Name: "Dummy",
		Age:  20,
	}

	if err := db.Create(dummy).Error; err != nil {
		panic(err)
	}
}
