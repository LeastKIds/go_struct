package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/model"
)

var Models = []interface{}{
	&model.Dummy{},
}

func main() {
	db, err := connection.NewConnection().Connect()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(Models...); err != nil {
		panic(err)
	}
}
