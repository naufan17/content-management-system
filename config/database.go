package config

import (
	"log"
	"strconv"

	// "github.com/naufan17/content-management-system/database/seeders"
	// "github.com/naufan17/content-management-system/internal/auth"
	// "github.com/naufan17/content-management-system/internal/category"
	// "github.com/naufan17/content-management-system/internal/comment"
	// "github.com/naufan17/content-management-system/internal/news"
	// "github.com/naufan17/content-management-system/internal/page"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = ConnectDB()
}

func ConnectDB() *gorm.DB {
	cfg := LoadConfig()

	dbHost := cfg.DBHost
	dbUser := cfg.DBUser
	dbPassword := cfg.DBPassword
	dbName := cfg.DBName
	dbPort := cfg.DBPort
	dbSsl := cfg.DBSsl
	dbTimezone := cfg.DBTimezone

	dbMaxIdle, err := strconv.Atoi(cfg.DBMaxIdle)

	if err != nil {
		log.Fatal("Invalid DBMaxIdle value:", err)
	}

	dbMaxOpen, err := strconv.Atoi(cfg.DBMaxOpen)

	if err != nil {
		log.Fatal("Invalid DBMaxOpen value:", err)
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSsl + " TimeZone=" + dbTimezone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get SQL DB from GORM:", err)
	}

	sqlDB.SetMaxIdleConns(dbMaxIdle)
	sqlDB.SetMaxOpenConns(dbMaxOpen)

	log.Println("Connected to database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&auth.User{},
		&category.Category{},
		&page.Page{},
		&news.News{},
		&comment.Comment{},
	)

// 	if err != nil {
// 		log.Fatal("Failed to migrate database", err)
// 	} else {
// 		log.Println("Database migrated successfully")
// 	}
// }

// func SeedAll(db *gorm.DB) {
// 	seeders.SeedUsers(db)
// 	seeders.SeedCategories(db)

// 	log.Println("Database seeded successfully")
// }
