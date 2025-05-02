package repositories

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryByID(id uuid.UUID) (models.Category, error) {
	var category models.Category

	if err := config.DB.Where("id = ?", id).Select("id", "name").First(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(id uuid.UUID, category models.Category) (models.Category, error) {
	if err := config.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func DeleteCategory(id uuid.UUID) error {
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
