package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/internal/handlers"
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
