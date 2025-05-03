package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/pkg/config"
	"github.com/naufan17/content-management-system/pkg/utils"
)

type Handler struct {
	pageService PageService
}

func NewHandler(pageService PageService) *Handler {
	return &Handler{pageService: pageService}
}

func (h *Handler) GetPages(c *gin.Context) {
	pages, err := h.pageService.GetPages()

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

func (h *Handler) GetPage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	page, err := h.pageService.GetPage(uuidID)

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

func (h *Handler) CreatePage(c *gin.Context) {
	var page CreatePageRequest

	err := c.ShouldBindJSON(&page)

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

	claimUser := c.MustGet("claimsUser").(*utils.Claims)
	page.UserID = claimUser.Sub
	err = h.pageService.CreatePage(page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create page",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "page created successfully",
	})
}

func (h *Handler) UpdatePage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	var page UpdatePageRequest

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

	userID := c.MustGet("claimsUser").(*utils.Claims).Sub.String()
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse user ID",
		})

		return
	}

	page.UserID = parsedUserID
	err = h.pageService.UpdatePage(uuidID, page)

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

func (h *Handler) DeletePage(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid page ID format",
		})

		return
	}

	userID := c.MustGet("claimsUser").(*utils.Claims).Sub.String()
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse user ID",
		})

		return
	}

	err = h.pageService.DeletePage(uuidID, parsedUserID)

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
