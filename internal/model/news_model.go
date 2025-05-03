package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID         uuid.UUID      `json:"id" gorm:"type:char(36);not null"`
	UserID     uuid.UUID      `json:"user_id" gorm:"type:char(36);not null"`
	CategoryID uuid.UUID      `json:"category_id" gorm:"type:char(36);not null"`
	Title      string         `json:"title" gorm:"type:varchar(255);not null"`
	Content    string         `json:"content" gorm:"type:text;not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (news *News) BeforeCreate(tx *gorm.DB) (err error) {
	news.ID = uuid.New()

	return nil
}
