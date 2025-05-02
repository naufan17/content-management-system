package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/config"
	"github.com/naufan17/content-management-system/internal/dtos"
	"github.com/naufan17/content-management-system/internal/services"
	"github.com/naufan17/content-management-system/pkg/util"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "categories not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get categories",
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

	category, err := services.GetCategory(uuidID)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "category not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get category",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	var category dtos.CreateCategoryDto

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(category); validatorErr != nil {
		errors := util.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	_, err := services.CreateCategory(category)

	if err != nil {
		if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create category",
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
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

	var category dtos.UpdateCategoryDto

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(category); validatorErr != nil {
		errors := util.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	_, err = services.UpdateCategory(uuidID, category)

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "record not found" {
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

	err = services.DeleteCategory(uuidID)

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "record not found" {
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
