package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/pkg/auth"
)

func AuthorizeBearer(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header is required",
		})

		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := auth.ValidateJWTAccess(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})

		c.Abort()
		return
	}

	c.Set("claimsUser", claims)
	c.Next()
}
