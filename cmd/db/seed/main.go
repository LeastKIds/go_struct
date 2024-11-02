package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

func main() {
	db, err := connection.NewConnection().Connect()
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
