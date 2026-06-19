package server

import (
	"go-conceptual/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RoutesHandler(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.HTTPResponse{
			Success: true,
			Message: "Server is Running Successfully!",
		})
	})

}
