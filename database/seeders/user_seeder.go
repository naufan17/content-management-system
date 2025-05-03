package seeders

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/naufan17/content-management-system/internal/model"
)

func SeedUsers(db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Password1234"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	users := []model.User{
		{Name: "John Doe", Username: "jhon", Password: string(hashedPassword)},
		{Name: "Jane Doe", Username: "jane", Password: string(hashedPassword)},
		{Name: "Mark Doe", Username: "mark", Password: string(hashedPassword)},
		{Name: "Alice Doe", Username: "alice", Password: string(hashedPassword)},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
}
