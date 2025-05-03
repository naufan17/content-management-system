package repository

import (
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/model"
)

func FindAllPage() ([]model.Page, error) {
	var pages []model.Page

	err := config.DB.Find(&pages).Error

	return pages, err
}

func FindByIDPage(id uuid.UUID) (model.Page, error) {
	var page model.Page

	err := config.DB.Where("id = ?", id).First(&page).Error

	return page, err
}

func CreatePage(page model.Page) error {
	return config.DB.Create(&page).Error
}

func UpdatePage(id uuid.UUID, page model.Page) error {
	return config.DB.Model(&model.Page{}).Where("id = ?", id).Updates(&page).Error
}

func DeletePage(id uuid.UUID) error {
	return config.DB.Delete(&model.Page{}, id).Error
}
