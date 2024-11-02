package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LeastKIds/go_struct/internal/infrastructure/config"
)

const (
	developDatabaseURL = "user:password@tcp(localhost:3306)/go_struct?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	content, err := os.ReadFile(".env.example")
	if err != nil {
		fmt.Println("read file err:", err)
		panic(err)
	}

	// 개발 환경의 env 파일을 생성합니다.
	envContent := string(content)
	envContent = strings.Replace(envContent, "DATABASE_URL=", fmt.Sprintf(`DATABASE_URL="%s"`, developDatabaseURL), 1)
	envContent = strings.Replace(envContent, "ENV=", fmt.Sprintf(`ENV="%s"`, config.EnvDevelopment), 1)
	if err := os.WriteFile(".env", []byte(envContent), 0o644); err != nil {
		fmt.Println("write file err:", err)
		panic(err)
	}

}
