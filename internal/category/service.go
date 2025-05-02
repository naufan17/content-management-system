package category

import (
	"errors"

	"github.com/google/uuid"
)

type CategoryService interface {
	GetCategories() ([]CategoryDto, error)
	GetCategory(id uuid.UUID) (CategoryDto, error)
	CreateCategory(category CreateCategoryDto) error
	UpdateCategory(id uuid.UUID, category UpdateCategoryDto) error
	DeleteCategory(id uuid.UUID) error
}

type categoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(categoryRepository CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) GetCategories() ([]CategoryDto, error) {
	categories, err := s.categoryRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var categoryDtos []CategoryDto

	for _, category := range categories {
		categoryDtos = append(categoryDtos, CategoryModelToDto(category))
	}

	return categoryDtos, nil
}

func (s *categoryService) GetCategory(id uuid.UUID) (CategoryDto, error) {
	category, err := s.categoryRepository.FindByID(id)

	if err != nil {
		return CategoryDto{}, err
	}

	return CategoryModelToDto(category), nil
}

func (s *categoryService) CreateCategory(category CreateCategoryDto) error {
	err := s.categoryRepository.Create(CreateCategoryDtoToModel(category))

	if err != nil {
		return err
	}

	return nil
}

func (s *categoryService) UpdateCategory(id uuid.UUID, category UpdateCategoryDto) error {
	_, err := s.categoryRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	err = s.categoryRepository.Update(id, UpdateCategoryDtoToModel(category))

	if err != nil {
		return err
	}

	return nil
}

func (s *categoryService) DeleteCategory(id uuid.UUID) error {
	_, err := s.categoryRepository.FindByID(id)

	if err != nil {
		return errors.New("not found")
	}

	err = s.categoryRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
