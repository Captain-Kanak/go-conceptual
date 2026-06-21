package server

import (
	"go-conceptual/internal/domain/user"
	"go-conceptual/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RoutesHandler(db *gorm.DB, e *echo.Echo) {
	db.AutoMigrate(user.User{})

	// health check
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is Running Successfully!",
		})
	})

	// user routes
	user.Routes(db, e)
}
