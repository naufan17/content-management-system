package news

import (
	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/pkg/middleware"
)

func NewsRoute(r *gin.Engine, h *Handler) {
	api := r.Group("/api/news")
	{
		api.GET("/", h.GetNews)
		api.GET("/:id", h.GetNewsByID)
		api.POST("/", middleware.AuthorizeBearer, h.CreateNews)
		api.PUT("/:id", middleware.AuthorizeBearer, h.UpdateNews)
		api.DELETE("/:id", middleware.AuthorizeBearer, h.DeleteNews)
	}
}
