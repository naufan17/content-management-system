package page

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PageRepository interface {
	FindAll() ([]Page, error)
	FindByID(id uuid.UUID) (Page, error)
	Create(page Page) error
	Update(id uuid.UUID, page Page) error
	Delete(id uuid.UUID) error
}

type pageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return &pageRepository{db: db}
}

func (r *pageRepository) FindAll() ([]Page, error) {
	var pages []Page

	err := r.db.Find(&pages).Error

	return pages, err
}

func (r *pageRepository) FindByID(id uuid.UUID) (Page, error) {
	var page Page

	err := r.db.Where("id = ?", id).First(&page).Error

	return page, err
}

func (r *pageRepository) Create(page Page) error {
	return r.db.Create(&page).Error
}

func (r *pageRepository) Update(id uuid.UUID, page Page) error {
	return r.db.Model(&Page{}).Where("id = ?", id).Updates(&page).Error
}

func (r *pageRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Page{}, id).Error
}
