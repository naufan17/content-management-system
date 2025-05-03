package category

import (
	"github.com/google/uuid"
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

func CategoryModelToDto(category Category) CategoryDto {
	return CategoryDto{
		ID:   category.ID,
		Name: category.Name,
	}
}

func CreateCategoryDtoToModel(category CreateCategoryRequest) Category {
	return Category{
		Name: category.Name,
	}
}

func UpdateCategoryDtoToModel(category UpdateCategoryRequest) Category {
	return Category{
		Name: category.Name,
	}
}
