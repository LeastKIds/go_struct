package server

import (
	"net/http"

	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
	"github.com/LeastKIds/go_struct/internal/infrastructure/database"
	"github.com/LeastKIds/go_struct/internal/infrastructure/framework"
	"gorm.io/gorm"
)

func Run() {
	cfg := config.NewConfig()
	db, err := database.NewConnection(cfg.DATABASE_URL).Connect()
	if err != nil {
		panic(err)
	}
	f := framework.NewFramework(db)
	v1 := f.Group("/v1")
	v1.GET("/dummy", func(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
		w.Write([]byte("Hello, World!"))
	})

	f.Start(":8080")
}
