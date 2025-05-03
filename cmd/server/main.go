package main

import (
	"github.com/naufan17/content-management-system/internal/auth"
	"github.com/naufan17/content-management-system/internal/category"
	"github.com/naufan17/content-management-system/internal/news"
	"github.com/naufan17/content-management-system/internal/page"
	"github.com/naufan17/content-management-system/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDB()
	env := cfg.GinMode
	port := cfg.Port
	router := gin.Default()

	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	config.SetupHelmet(router)
	config.SetupCORS(router)
	config.SetupRateLimit(router)

	// Auth module
	authRepo := auth.NewUserRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewHandler(authService)
	auth.AuthRoute(router, authHandler)

	// Category module
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)
	category.CategoryRoute(router, categoryHandler)

	// News module
	newsRepo := news.NewNewsRepository(db)
	newsService := news.NewNewsService(newsRepo)
	newsHandler := news.NewHandler(newsService)
	news.NewsRoute(router, newsHandler)

	// Page module
	pageRepo := page.NewPageRepository(db)
	pageService := page.NewPageService(pageRepo)
	pageHandler := page.NewHandler(pageService)
	page.PageRoute(router, pageHandler)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":" + port)
}
