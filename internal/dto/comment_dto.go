package dto

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/internal/model"
)

type CommentDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	CreatedAt string    `json:"created_at"`
}

type CreateCommentRequest struct {
	Name    string    `json:"name"`
	Comment string    `json:"comment" validate:"required,max=255"`
	NewsID  uuid.UUID `json:"news_id"`
}

func CommentModelToDto(comment model.Comment) CommentDto {
	return CommentDto{
		ID:        comment.ID,
		Name:      comment.Name,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt.String(),
	}
}

func CreateCommentDtoToModel(comment CreateCommentRequest) model.Comment {
	return model.Comment{
		Name:    comment.Name,
		Comment: comment.Comment,
		NewsID:  comment.NewsID,
	}
}
