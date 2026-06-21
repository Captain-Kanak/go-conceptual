package user

import (
	"go-conceptual/internal/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, e *echo.Echo) {
	repo := NewRepository(db)
	jwt := auth.NewJWTService("", 0)
	service := NewService(repo, jwt)
	handler := NewHandler(service)

	api := e.Group("/api/v1")

	api.POST("/auth/register", handler.RegisterUser)
	api.POST("/auth/login", handler.LoginUser)
	api.GET("/users", handler.GetAllUsers)
}
