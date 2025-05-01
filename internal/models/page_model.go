package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	ID          uuid.UUID      `json:"id" gorm:"type:char(36);not null"`
	UserID      uuid.UUID      `json:"user_id" gorm:"type:char(36);not null"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	CustomURL   string         `json:"custom_url" gorm:"type:varchar(255);not null"`
	Content     string         `json:"content" gorm:"type:text;not null"`
	IsPublished bool           `json:"is_published" gorm:"default:true;not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (page *Page) BeforeCreate(tx *gorm.DB) (err error) {
	page.ID = uuid.New()

	return nil
}
