package category

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	FindAll() ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
	Create(category Category) error
	Update(id uuid.UUID, category Category) error
	Delete(id uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() ([]Category, error) {
	var categories []Category

	err := r.db.Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) FindByID(id uuid.UUID) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", id).First(&category).Error

	return category, err
}

func (r *categoryRepository) Create(category Category) error {
	return r.db.Create(&category).Error
}

func (r *categoryRepository) Update(id uuid.UUID, category Category) error {
	return r.db.Model(&Category{}).Where("id = ?", id).Updates(&category).Error
}

func (r *categoryRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Category{}, id).Error
}
