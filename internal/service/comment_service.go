package service

import (
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/repository"
)

func GetCommentsByNewsId(newsId uuid.UUID) ([]dto.CommentDto, error) {
	comments, err := repository.FindAllCommentsByNewsID(newsId)

	if err != nil {
		return nil, err
	}

	var commentDtos []dto.CommentDto

	for _, comment := range comments {
		commentDtos = append(commentDtos, dto.CommentModelToDto(comment))
	}

	return commentDtos, nil
}

func CreateComment(comment dto.CreateCommentRequest, userID *uuid.UUID) error {
	var userName string

	if userID != nil {
		user, err := repository.FindByIDUser(*userID)

		if err != nil {
			return err
		}

		userName = user.Name
	} else {
		userName = "anonymous"
	}

	comment.Name = userName
	commentModel := dto.CreateCommentDtoToModel(comment)
	err := repository.CreateComment(commentModel)

	if err != nil {
		return err
	}

	return nil
}
