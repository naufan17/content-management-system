package repository

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/model"
)

func CreateUser(user model.User) error {
	return config.DB.Create(&user).Error
}

func FindByUsernameUser(username string) (model.User, error) {
	var user model.User

	err := config.DB.Where("username = ?", username).First(&user).Error

	return user, err
}

func FindByIDUser(id uuid.UUID) (model.User, error) {
	var user model.User

	err := config.DB.Where("id = ?", id).First(&user).Error

	return user, err
}
