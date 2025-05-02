package category

import (
	"github.com/google/uuid"
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

func CategoryModelToDto(category Category) CategoryDto {
	return CategoryDto{
		ID:   category.ID,
		Name: category.Name,
	}
}

func CreateCategoryDtoToModel(category CreateCategoryDto) Category {
	return Category{
		Name: category.Name,
	}
}

func UpdateCategoryDtoToModel(category UpdateCategoryDto) Category {
	return Category{
		Name: category.Name,
	}
}
