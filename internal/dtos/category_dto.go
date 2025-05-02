package dtos

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/internal/models"
)

type CategoryDto struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateCategoryDto struct {
	Name string `json:"name" validate:"required,max=50"`
}

type UpdateCategoryDto struct {
	Name string `json:"name" validate:"required,max=50"`
}

func CategoryModelToDto(category models.Category) CategoryDto {
	return CategoryDto{
		ID:   category.ID,
		Name: category.Name,
	}
}

func CreateCategoryDtoToModel(category CreateCategoryDto) models.Category {
	return models.Category{
		Name: category.Name,
	}
}

func UpdateCategoryDtoToModel(category UpdateCategoryDto) models.Category {
	return models.Category{
		Name: category.Name,
	}
}
