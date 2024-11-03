package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodGet {
				fmt.Println("GET transaction")
				c.Set("db", db)
				return next(c)
			}

			fmt.Println("POST transaction")

			tx := db.Begin()
			if tx.Error != nil {
				return tx.Error
			}

			c.Set("db", tx)
			err := next(c)

			if err != nil {
				tx.Rollback()
				return err
			}

			return tx.Commit().Error
		}
	}
}
