package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/naufan17/content-management-system/internal/handler"
	"github.com/naufan17/content-management-system/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handler.Login)
		}

		category := api.Group("/categories")
		{
			category.GET("/", handler.GetCategories)
			category.GET("/:id", handler.GetCategory)
			category.POST("/", middleware.AuthorizeBearer, handler.CreateCategory)
			category.PUT("/:id", middleware.AuthorizeBearer, handler.UpdateCategory)
			category.DELETE("/:id", middleware.AuthorizeBearer, handler.DeleteCategory)
		}

		news := api.Group("/news")
		{
			news.GET("/", handler.GetNews)
			news.GET("/:id", handler.GetNewsByID)
			news.POST("/", middleware.AuthorizeBearer, handler.CreateNews)
			news.PUT("/:id", middleware.AuthorizeBearer, handler.UpdateNews)
			news.DELETE("/:id", middleware.AuthorizeBearer, handler.DeleteNews)
		}

		page := api.Group("/pages")
		{
			page.GET("/", handler.GetPages)
			page.GET("/:id", handler.GetPage)
			page.POST("/", middleware.AuthorizeBearer, handler.CreatePage)
			page.PUT("/:id", middleware.AuthorizeBearer, handler.UpdatePage)
			page.DELETE("/:id", middleware.AuthorizeBearer, handler.DeletePage)
		}
	}
}
