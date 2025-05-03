package page

import (
	"github.com/google/uuid"
)

type PageDto struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	CustomURL   string    `json:"custom_url"`
	Content     string    `json:"content"`
	IsPublished bool      `json:"is_published"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   string    `json:"created_at"`
}

type CreatePageRequest struct {
	Title       string    `json:"title" validate:"required,max=100"`
	Content     string    `json:"content" validate:"required"`
	UserID      uuid.UUID `json:"user_id"`
	CustomURL   string    `json:"custom_url" validate:"required,max=100,url"`
	IsPublished bool      `json:"is_published" validate:"required" default:"true"`
}

type UpdatePageRequest struct {
	Title       string    `json:"title" validate:"required,max=100"`
	Content     string    `json:"content" validate:"required"`
	UserID      uuid.UUID `json:"user_id"`
	CustomURL   string    `json:"custom_url" validate:"required,max=100,url"`
	IsPublished bool      `json:"is_published" validate:"required" default:"true"`
}

func PageModelToDto(page Page) PageDto {
	return PageDto{
		ID:          page.ID,
		Title:       page.Title,
		CustomURL:   page.CustomURL,
		Content:     page.Content,
		IsPublished: page.IsPublished,
		UserID:      page.UserID,
		CreatedAt:   page.CreatedAt.String(),
	}
}

func CreatePageDtoToModel(page CreatePageRequest) Page {
	return Page{
		Title:       page.Title,
		Content:     page.Content,
		UserID:      page.UserID,
		CustomURL:   page.CustomURL,
		IsPublished: page.IsPublished,
	}
}

func UpdatePageDtoToModel(page UpdatePageRequest) Page {
	return Page{
		Title:       page.Title,
		Content:     page.Content,
		UserID:      page.UserID,
		CustomURL:   page.CustomURL,
		IsPublished: page.IsPublished,
	}
}
