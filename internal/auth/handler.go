package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/naufan17/content-management-system/pkg/config"
	"github.com/naufan17/content-management-system/pkg/utils"
)

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(user); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	accessToken, err := h.authService.LoginUser(user)

	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "username or password is incorrect",
			})

			return
		} else if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to login user",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to login user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": accessToken.AccessToken,
			"token_type":   accessToken.TokenType,
			"expires_in":   accessToken.ExpiresIn,
		},
	})
}
