package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Email       string         `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string         `json:"password" gorm:"type:varchar(255);not null"`
	DateOfBirth time.Time      `json:"age" gorm:"type:timestamp"`
	Phone       string         `json:"phone" gorm:"type:varchar(20)"`
	Address     string         `json:"address" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"type:timestamp;index"`
}
