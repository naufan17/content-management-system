package repository

import (
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/model"
)

func FindAllNews() ([]model.News, error) {
	var news []model.News

	err := config.DB.Find(&news).Error

	return news, err
}

func FindByIDNews(id uuid.UUID) (model.News, error) {
	var news model.News

	err := config.DB.Where("id = ?", id).First(&news).Error

	return news, err
}

func CreateNews(news model.News) error {
	return config.DB.Create(&news).Error
}

func UpdateNews(id uuid.UUID, news model.News) error {
	return config.DB.Model(&model.News{}).Where("id = ?", id).Updates(&news).Error
}

func DeleteNews(id uuid.UUID) error {
	return config.DB.Delete(&model.News{}, id).Error
}
