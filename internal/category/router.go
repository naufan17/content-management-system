package category

import (
	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/pkg/middleware"
)

func CategoryRoute(r *gin.Engine, h *Handler) {
	api := r.Group("/api/categories")
	{
		api.GET("/", h.GetCategories)
		api.GET("/:id", h.GetCategory)
		api.POST("/", middleware.AuthorizeBearer(), h.CreateCategory)
		api.PUT("/:id", middleware.AuthorizeBearer(), h.UpdateCategory)
		api.DELETE("/:id", middleware.AuthorizeBearer(), h.DeleteCategory)
	}
}
