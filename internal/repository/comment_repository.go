package repository

import (
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/model"
)

func FindAllCommentsByNewsID(newsID uuid.UUID) ([]model.Comment, error) {
	var comments []model.Comment

	err := config.DB.Where("news_id = ?", newsID).Find(&comments).Error

	return comments, err
}

func CreateComment(comment model.Comment) error {
	return config.DB.Create(&comment).Error
}
