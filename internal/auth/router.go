package auth

import "github.com/gin-gonic/gin"

func AuthRoute(r *gin.Engine, h *Handler) {
	api := r.Group("/api/auth")
	{
		api.POST("/login", h.Login)
	}
}
