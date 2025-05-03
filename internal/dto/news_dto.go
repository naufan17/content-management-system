package dto

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/internal/model"
)

type NewsDto struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`
	CreatedAt  string    `json:"created_at"`
}

type CreateNewsRequest struct {
	Title      string    `json:"title" validate:"required,max=100"`
	Content    string    `json:"content" validate:"required"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id" validate:"required"`
}

type UpdateNewsRequest struct {
	Title      string    `json:"title" validate:"required,max=100"`
	Content    string    `json:"content" validate:"required"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id" validate:"required"`
}

func NewsModelToDto(news model.News) NewsDto {
	return NewsDto{
		ID:         news.ID,
		Title:      news.Title,
		Content:    news.Content,
		UserID:     news.UserID,
		CategoryID: news.CategoryID,
		CreatedAt:  news.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func CreateNewsDtoToModel(news CreateNewsRequest) model.News {
	return model.News{
		Title:      news.Title,
		Content:    news.Content,
		CategoryID: news.CategoryID,
		UserID:     news.UserID,
	}
}

func UpdateNewsDtoToModel(news UpdateNewsRequest) model.News {
	return model.News{
		Title:      news.Title,
		Content:    news.Content,
		CategoryID: news.CategoryID,
		UserID:     news.UserID,
	}
}
