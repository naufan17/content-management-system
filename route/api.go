package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/internal/handlers"
	"github.com/naufan17/content-management-system/internal/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/naufan17/content-management-system/docs"
)

func ApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", handlers.GetCategories)
			categories.GET("/:id", handlers.GetCategory)
			categories.POST("/", middlewares.AuthorizeBearer, handlers.CreateCategory)
			categories.PUT("/:id", middlewares.AuthorizeBearer, handlers.UpdateCategory)
			categories.DELETE("/:id", middlewares.AuthorizeBearer, handlers.DeleteCategory)
		}

		// news := api.Group("/news")
		// {
		// 	news.GET("/", handlers.GetNews)
		// 	news.GET("/:id", handlers.GetNewsByID)
		// news.POST("/", middlewares.AuthorizeBearer(), handlers.CreateNews)
		// news.PUT("/:id", middlewares.AuthorizeBearer(), handlers.UpdateNews)
		// news.DELETE("/:id", middlewares.AuthorizeBearer(), handlers.DeleteNews)

		// 	news.POST("/:id/comments", handlers.CreateComment)
		// }

		// pages := api.Group("/pages")
		// {
		// 	pages.GET("/", handlers.GetPages)
		// 	pages.GET("/:id", handlers.GetPageByID)
		// pages.POST("/", middlewares.AuthorizeBearer(), handlers.CreatePage)
		// pages.PUT("/:id", middlewares.AuthorizeBearer(), handlers.UpdatePage)
		// pages.DELETE("/:id", middlewares.AuthorizeBearer(), handlers.DeletePage)
		// }
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "method not allowed",
		})
	})
}
