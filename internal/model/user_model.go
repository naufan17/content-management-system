package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);not null"`
	Name      string         `json:"name" gorm:"type:varchar(128);not null"`
	Username  string         `json:"username" gorm:"type:varchar(128);not null;unique"`
	Password  string         `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	return nil
}
