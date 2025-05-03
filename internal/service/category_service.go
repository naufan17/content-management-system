package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/repository"
)

func GetCategories() ([]dto.CategoryDto, error) {
	categories, err := repository.FindAllCategory()

	if err != nil {
		return nil, err
	}

	var categoryDtos []dto.CategoryDto

	for _, category := range categories {
		categoryDtos = append(categoryDtos, dto.CategoryModelToDto(category))
	}

	return categoryDtos, nil
}

func GetCategory(id uuid.UUID) (dto.CategoryDto, error) {
	category, err := repository.FindByIDCategory(id)

	if err != nil {
		return dto.CategoryDto{}, err
	}

	return dto.CategoryModelToDto(category), nil
}

func CreateCategory(category dto.CreateCategoryRequest) error {
	err := repository.CreateCategory(dto.CreateCategoryDtoToModel(category))

	if err != nil {
		return err
	}

	return nil
}

func UpdateCategory(id uuid.UUID, category dto.UpdateCategoryRequest) error {
	_, err := repository.FindByIDCategory(id)

	if err != nil {
		return errors.New("not found")
	}

	err = repository.UpdateCategory(id, dto.UpdateCategoryDtoToModel(category))

	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(id uuid.UUID) error {
	_, err := repository.FindByIDCategory(id)

	if err != nil {
		return errors.New("not found")
	}

	err = repository.DeleteCategory(id)

	if err != nil {
		return err
	}

	return nil
}
