package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/naufan17/content-management-system/pkg/config"
	"github.com/naufan17/content-management-system/pkg/utils"
)

type Handler struct {
	newsService NewsService
}

func NewHandler(newsService NewsService) *Handler {
	return &Handler{
		newsService: newsService,
	}
}

func (h *Handler) GetNews(c *gin.Context) {
	news, err := h.newsService.GetNews()

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

func (h *Handler) GetNewsByID(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	news, err := h.newsService.GetNewsByID(uuidID)

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

func (h *Handler) CreateNews(c *gin.Context) {
	var news CreateNewsRequest

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

	claimUser := c.MustGet("claimsUser").(*utils.Claims)
	userID := claimUser.Sub.String()
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse user ID",
		})

		return
	}

	news.UserID = parsedUserID
	err = h.newsService.CreateNews(news)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create news",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "news created successfully",
	})
}

func (h *Handler) UpdateNews(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
		})

		return
	}

	var news UpdateNewsRequest

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

	userID := c.MustGet("claimsUser").(*utils.Claims).Sub.String()
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse user ID",
		})

		return
	}

	news.UserID = parsedUserID
	err = h.newsService.UpdateNews(uuidID, news)

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

func (h *Handler) DeleteNews(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid news ID format",
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

	err = h.newsService.DeleteNews(uuidID, parsedUserID)

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
