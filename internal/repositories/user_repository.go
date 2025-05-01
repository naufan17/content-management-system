package repositories

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/models"
)

func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := config.DB.Where("username = ?", username).Select("id", "username", "password").First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).Select("id", "name", "username").First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
