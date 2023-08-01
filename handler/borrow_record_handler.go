package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ericjovian/gin-template/dto"
	"github.com/ericjovian/gin-template/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBorrowRecord(c *gin.Context) {
	var req dto.BorrowRecordRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.borrowRecordUsecase.CreateBorrowRecord(req)
	if err != nil {
		if errors.Is(err, utils.ErrUserNotFound) || errors.Is(err, utils.ErrBookNotFound) || errors.Is(err, utils.ErrEmptyBook) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    utils.ErrCodeBadRequest,
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

func (h *Handler) ReturnBorrowRecord(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		})
		return
	}

	res, err := h.borrowRecordUsecase.ReturnBook(intId)
	if err != nil {
		if errors.Is(err, utils.ErrRecordNotFound) || errors.Is(err, utils.ErrAlreadyReturned) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    utils.ErrCodeBadRequest,
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

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
