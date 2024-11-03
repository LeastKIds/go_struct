package server

import (
	"net/http"

	"github.com/LeastKIds/go_struct/internal/infrastructure/framework"
)

func Run() {
	f := framework.NewFramework()
	v1 := f.Group("/v1")
	v1.GET("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	f.Start(":8080")
}
