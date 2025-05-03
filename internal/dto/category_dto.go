package dto

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/internal/model"
)

type CategoryDto struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}

func CategoryModelToDto(category model.Category) CategoryDto {
	return CategoryDto{
		ID:   category.ID,
		Name: category.Name,
	}
}

func CreateCategoryDtoToModel(category CreateCategoryRequest) model.Category {
	return model.Category{
		Name: category.Name,
	}
}

func UpdateCategoryDtoToModel(category UpdateCategoryRequest) model.Category {
	return model.Category{
		Name: category.Name,
	}
}
