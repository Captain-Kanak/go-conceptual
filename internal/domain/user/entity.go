package user

import (
	"go-conceptual/internal/domain/user/dto"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Email       string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string         `json:"-" gorm:"type:varchar(255);not null"`
	DateOfBirth *time.Time     `json:"date_of_birth" gorm:"type:timestamp"`
	Phone       string         `json:"phone" gorm:"type:varchar(20)"`
	Address     string         `json:"address" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"type:timestamp;index"`
}

func (u *User) ToResponse() *dto.Response {
	return &dto.Response{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		DateOfBirth: u.DateOfBirth,
		Phone:       u.Phone,
		Address:     u.Address,
		CreatedAt:   u.CreatedAt,
	}
}

func (u *User) hashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) checkPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
