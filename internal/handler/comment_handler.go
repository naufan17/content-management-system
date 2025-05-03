package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/service"
	"github.com/naufan17/content-management-system/pkg/utils"
)

func GetCommentsByNewsID(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	comments, err := service.GetCommentsByNewsId(uuidID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "news not found",
		})

		return
	} else if len(comments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "news not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func CreateComment(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	var comment dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(comment); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	claimsUser, exists := c.Get("claimsUser")
	var userID string
	var userUUID *uuid.UUID

	if exists {
		userID = claimsUser.(*utils.Claims).Sub.String()
		parsedUUID, parseErr := uuid.Parse(userID)

		if parseErr == nil {
			userUUID = &parsedUUID
		}
	}

	comment.NewsID = uuidID
	err = service.CreateComment(comment, userUUID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create comment",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment created successfully",
	})
}
