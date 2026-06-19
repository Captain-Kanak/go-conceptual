package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(u User) error
	GetAll() ([]User, error)
	GetByID(id uuid.UUID) (User, error)
	// UpdateByID(id uuid.UUID, u User) error
	// DeleteByID(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(u User) error {
	return r.db.Create(&u).Error
}

func (r *repository) GetAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetByID(id uuid.UUID) (User, error) {
	var u User

	err := r.db.First(&u, id).Error

	return u, err
}

func (r *repository) UpdateByID(id uuid.UUID, u User) error {
	return r.db.Model(&User{}).Where("id = ?", id).Updates(u).Error
}

func (r *repository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&User{}, id).Error
}
