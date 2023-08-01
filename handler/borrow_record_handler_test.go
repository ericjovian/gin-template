package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"git.garena.com/sea-labs-id/batch-05/gin-template/dto"
	mocks "git.garena.com/sea-labs-id/batch-05/gin-template/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/gin-template/server"
	"git.garena.com/sea-labs-id/batch-05/gin-template/testutils"
	"git.garena.com/sea-labs-id/batch-05/gin-template/utils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestCreateBorrowRecord(t *testing.T) {
	reqBorrow := dto.BorrowRecordRequest{
		UserId: 1,
		BookId: 1,
	}
	resBorrow := dto.BorrowRecordResponse{
		ID:         1,
		UserId:     1,
		BookId:     1,
		Status:     "borrow",
		BorrowDate: time.Now(),
	}

	t.Run("should return borrow record when status code 201", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"data": resBorrow,
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("CreateBorrowRecord", reqBorrow).Return(&resBorrow, nil)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(reqBorrow)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when body request is invalid", func(t *testing.T) {
		invalidRequest := dto.BorrowRecordRequest{}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidRequest)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when user not found", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrUserNotFound.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("CreateBorrowRecord", reqBorrow).Return(nil, utils.ErrUserNotFound)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(reqBorrow)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when book not found", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrBookNotFound.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("CreateBorrowRecord", reqBorrow).Return(nil, utils.ErrBookNotFound)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(reqBorrow)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when book's stock is empty", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrEmptyBook.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("CreateBorrowRecord", reqBorrow).Return(nil, utils.ErrEmptyBook)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(reqBorrow)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code INTERNAL_SERVER_ERROR when server error occured", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeInternalServerError,
			"message": "error",
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("CreateBorrowRecord", reqBorrow).Return(nil, errors.New("error"))
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(reqBorrow)

		req, _ := http.NewRequest(http.MethodPost, "/borrow-records", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})
}

func TestReturnBorrowRecord(t *testing.T) {
	id := 1
	now := time.Now()
	resBorrow := dto.BorrowRecordResponse{
		ID:         1,
		UserId:     1,
		BookId:     1,
		Status:     "borrow",
		BorrowDate: now,
		ReturnDate: &now,
	}

	t.Run("should return updated borrow record when status code 200", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"data": resBorrow,
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("ReturnBook", id).Return(&resBorrow, nil)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/borrow-records/%v", id), nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when id param is not an integer", func(t *testing.T) {
		invalidParam := "test"
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/borrow-records/%v", invalidParam), nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when record not found", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrRecordNotFound.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("ReturnBook", id).Return(nil, utils.ErrRecordNotFound)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/borrow-records/%v", id), nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when book already returned", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrAlreadyReturned.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("ReturnBook", id).Return(nil, utils.ErrAlreadyReturned)
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/borrow-records/%v", id), nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code INTERNAL_SERVER_ERROR when server error occured", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeInternalServerError,
			"message": "error",
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBorrowRecordUsecase(t)
		mockUsecase.On("ReturnBook", id).Return(nil, errors.New("error"))
		cfg := server.RouterConfig{
			BorrowRecordUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/borrow-records/%v", id), nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})
}
