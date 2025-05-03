package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/pkg/config"
	"github.com/naufan17/content-management-system/pkg/utils"
)

type Handler struct {
	categoryService CategoryService
}

func NewHandler(categoryService CategoryService) *Handler {
	return &Handler{categoryService: categoryService}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategories()

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

func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	category, err := h.categoryService.GetCategory(uuidID)

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

func (h *Handler) CreateCategory(c *gin.Context) {
	var category CreateCategoryRequest

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

	err := h.categoryService.CreateCategory(category)

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

func (h *Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	var category UpdateCategoryRequest

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

	err = h.categoryService.UpdateCategory(uuidID, category)

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

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category ID format",
		})

		return
	}

	err = h.categoryService.DeleteCategory(uuidID)

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
