package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);not null"`
	NewsID    uuid.UUID      `json:"news_id" gorm:"type:char(36);not null"`
	News      News           `json:"news" gorm:"foreignKey:NewsID;references:ID"`
	Name      string         `json:"name" gorm:"type:varchar(128);not null"`
	Comment   string         `json:"comment" gorm:"type:text;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	comment.ID = uuid.New()

	return nil
}
