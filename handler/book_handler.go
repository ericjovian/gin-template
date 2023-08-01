package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/gin-template/dto"
	"git.garena.com/sea-labs-id/batch-05/gin-template/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBooks(c *gin.Context) {
	res, err := h.bookUsecase.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    utils.ErrCodeInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (h *Handler) CreateBook(c *gin.Context) {
	var req dto.BookRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.bookUsecase.CreateBook(req)
	if err != nil {
		if err == utils.ErrDuplicateBook {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    utils.ErrCodeDuplicate,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    utils.ErrCodeInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": res,
	})
}
