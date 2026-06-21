package user

import (
	"go-conceptual/internal/domain/user/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetByEmail(email string) (*User, error)
	Create(u *User) error
	GetAll() ([]User, error)
	GetByID(id uuid.UUID) (User, error)
	UpdateByID(id uuid.UUID, u dto.UpdateRequest) error
	DeleteByID(id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo *repository) GetByEmail(email string) (*User, error) {
	var u User

	tx := repo.db.Where(&User{Email: email}).First(&u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &u, nil
}

func (r *repository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *repository) GetAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetByID(id uuid.UUID) (User, error) {
	var u User

	err := r.db.Where(&User{ID: id}).First(&u).Error

	return u, err
}

func (r *repository) UpdateByID(id uuid.UUID, u dto.UpdateRequest) error {
	result := r.db.Model(&User{}).Where("id = ?", id).Updates(u)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&User{}, id).Error
}
