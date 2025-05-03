package repository

import (
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/model"
)

func FindAllCategory() ([]model.Category, error) {
	var categories []model.Category

	err := config.DB.Find(&categories).Error

	return categories, err
}

func FindByIDCategory(id uuid.UUID) (model.Category, error) {
	var category model.Category

	err := config.DB.Where("id = ?", id).First(&category).Error

	return category, err
}

func CreateCategory(category model.Category) error {
	return config.DB.Create(&category).Error
}

func UpdateCategory(id uuid.UUID, category model.Category) error {
	return config.DB.Model(&model.Category{}).Where("id = ?", id).Updates(&category).Error
}

func DeleteCategory(id uuid.UUID) error {
	return config.DB.Delete(&model.Category{}, id).Error
}
