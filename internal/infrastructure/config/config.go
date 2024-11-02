package config

import (
	"fmt"

	"github.com/LeastKIds/go_struct/internal/infrastructure/database/connection"
)

const (
	EnvDevelopment = "development"
)

func Set() {
	_, err := connection.NewConnection().Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}
