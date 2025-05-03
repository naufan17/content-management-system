package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/content-management-system/pkg/utils"
)

func AuthorizeBearer(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header format must be Bearer {token}",
		})

		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateJWT(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})

		c.Abort()
		return
	}

	c.Set("claimsUser", claims)
	c.Next()
}
