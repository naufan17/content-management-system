package services

import (
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/internal/dtos"
	"github.com/naufan17/content-management-system/internal/repositories"
)

func GetCategories() ([]dtos.CategoryDto, error) {
	categories, err := repositories.GetAllCategories()

	if err != nil {
		return []dtos.CategoryDto{}, err
	}

	var categoryDtos []dtos.CategoryDto

	for _, category := range categories {
		categoryDtos = append(categoryDtos, dtos.CategoryModelToDto(category))
	}

	return categoryDtos, nil
}

func GetCategory(id uuid.UUID) (dtos.CategoryDto, error) {
	category, err := repositories.GetCategoryByID(id)

	if err != nil {
		return dtos.CategoryDto{}, err
	}

	return dtos.CategoryModelToDto(category), nil
}

func CreateCategory(category dtos.CreateCategoryDto) (dtos.CategoryDto, error) {
	newCategory, err := repositories.CreateCategory(dtos.CreateCategoryDtoToModel(category))

	if err != nil {
		return dtos.CategoryDto{}, err
	}

	return dtos.CategoryModelToDto(newCategory), nil
}

func UpdateCategory(id uuid.UUID, category dtos.UpdateCategoryDto) (dtos.CategoryDto, error) {
	updatedCategory, err := repositories.UpdateCategory(id, dtos.UpdateCategoryDtoToModel(category))

	if err != nil {
		return dtos.CategoryDto{}, err
	}

	return dtos.CategoryModelToDto(updatedCategory), nil
}

func DeleteCategory(id uuid.UUID) error {
	err := repositories.DeleteCategory(id)

	if err != nil {
		return err
	}

	return nil
}
