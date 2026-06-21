package dto

import "time"

type CreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateRequest struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"age"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
}
