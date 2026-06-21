package server

import (
	"go-conceptual/internal/domain/user"
	"go-conceptual/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RoutesHandler(e *echo.Echo, db *gorm.DB) {
	db.AutoMigrate(user.User{})

	// initial routes
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is Running Successfully!",
		})
	})

	// user routes
}
