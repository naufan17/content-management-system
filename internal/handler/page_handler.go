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

func GetPages(c *gin.Context) {
	pages, err := service.GetPages()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pages not found",
		})

		return
	} else if len(pages) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pages not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pages,
	})
}

func GetPage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	page, err := service.GetPage(uuidID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "page not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": page,
	})
}

func CreatePage(c *gin.Context) {
	var page dto.CreatePageRequest

	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(page); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	page.UserID = userID
	err := service.CreatePage(page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create page",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "page created successfully",
	})
}

func UpdatePage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	var page dto.UpdatePageRequest

	err = c.ShouldBindJSON(&page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(page); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	page.UserID = userID
	err = service.UpdatePage(uuidID, page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update page",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "page updated successfully",
	})
}

func DeletePage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	claimsUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimsUser.Sub

	err = service.DeletePage(uuidID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete page",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "page deleted successfully",
	})
}
