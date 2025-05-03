package seeders

import (
	"log"

	"github.com/naufan17/content-management-system/internal/category"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {

	categories := []category.Category{
		{Name: "Technology"},
		{Name: "Health"},
		{Name: "Lifestyle"},
		{Name: "Travel"},
		{Name: "Food"},
		{Name: "Education"},
		{Name: "Finance"},
		{Name: "Entertainment"},
		{Name: "Sports"},
		{Name: "Fashion"},
		{Name: "Science"},
		{Name: "Politics"},
		{Name: "Business"},
	}

	if err := db.Create(&categories).Error; err != nil {
		log.Fatalf("Failed to seed categories: %v", err)
	}
}
