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

func GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "categories not found",
		})

		return
	} else if len(categories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "categories not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	category, err := service.GetCategory(uuidID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	var category dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(category); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	err := service.CreateCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create category",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "category created successfully",
	})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	var category dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(category); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	err = service.UpdateCategory(uuidID, category)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "category not found",
			})

			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update category",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category updated successfully",
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	err = service.DeleteCategory(uuidID)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "category not found",
			})

			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete category",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}
