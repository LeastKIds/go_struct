package server

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
)

func Run() {
	cfg := config.NewConfig()
	db, err := database.NewConnection(cfg.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}

}
