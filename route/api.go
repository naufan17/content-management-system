package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/naufan17/content-management-system/internal/handler"
	"github.com/naufan17/content-management-system/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
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

			comment := news.Group("/:id/comments")
			{
				comment.GET("/", handler.GetCommentsByNewsID)
				comment.POST("/", middleware.AuthorizeOptionalBearer, handler.CreateComment)
			}
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "resource not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "method not allowed",
		})
	})
}
