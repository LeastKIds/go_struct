package config

import "github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"

func Set() {
	_, err := connection.NewConnection().Connect()
	if err != nil {
		panic(err)
	}
}
