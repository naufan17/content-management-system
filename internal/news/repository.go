package news

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsRepository interface {
	FindAll() ([]News, error)
	FindByID(id uuid.UUID) (News, error)
	Create(news News) error
	Update(id uuid.UUID, news News) error
	Delete(id uuid.UUID) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) FindAll() ([]News, error) {
	var news []News

	err := r.db.Find(&news).Error

	return news, err
}

func (r *newsRepository) FindByID(id uuid.UUID) (News, error) {
	var news News

	err := r.db.Where("id = ?", id).First(&news).Error

	return news, err
}

func (r *newsRepository) Create(news News) error {
	return r.db.Create(&news).Error
}

func (r *newsRepository) Update(id uuid.UUID, news News) error {
	return r.db.Model(&News{}).Where("id = ?", id).Updates(&news).Error
}

func (r *newsRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&News{}, id).Error
}
