package main

import "github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"

func main() {
	db, err := connection.NewConnection().Connect()
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
