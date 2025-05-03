package page

import (
	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/pkg/middleware"
)

func PageRoute(r *gin.Engine, h *Handler) {
	api := r.Group("/api/pages")
	{
		api.GET("/", h.GetPages)
		api.GET("/:id", h.GetPage)
		api.POST("/", middleware.AuthorizeBearer, h.CreatePage)
		api.PUT("/:id", middleware.AuthorizeBearer, h.UpdatePage)
		api.DELETE("/:id", middleware.AuthorizeBearer, h.DeletePage)
	}
}
