package auth

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user User) error
	FindByUsername(username string) (User, error)
	FindById(id uuid.UUID) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByUsername(username string) (User, error) {
	var user User

	err := r.db.Where("username = ?", username).First(&user).Error

	return user, err
}

func (r *userRepository) FindById(id uuid.UUID) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).First(&user).Error

	return user, err
}
