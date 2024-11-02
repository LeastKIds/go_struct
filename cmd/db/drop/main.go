package main

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
)

func main() {
	config := config.NewConfig()
	db, err := database.NewConnection(config.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

	tableNames, _ := db.Migrator().GetTables()

	for _, tableName := range tableNames {
		if err := db.Migrator().DropTable(tableName); err != nil {
			panic(err)
		}
	}
}
