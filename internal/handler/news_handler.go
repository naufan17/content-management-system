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

func GetNews(c *gin.Context) {
	news, err := service.GetNews()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "news not found",
		})

		return
	} else if len(news) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "news not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": news,
	})
}

func GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	news, err := service.GetNewsByID(uuidID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "news not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": news,
	})
}

func CreateNews(c *gin.Context) {
	var news dto.CreateNewsRequest

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(news); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	news.UserID = userID
	err := service.CreateNews(news)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create news",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "news created successfully",
	})
}

func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	var news dto.UpdateNewsRequest

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(news); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	news.UserID = userID
	err = service.UpdateNews(uuidID, news)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update news",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "news updated successfully",
	})
}

func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	err = service.DeleteNews(uuidID, userID)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "news not found",
			})

			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete news",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "news deleted successfully",
	})
}
