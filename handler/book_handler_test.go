package handler_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ericjovian/gin-template/dto"
	"github.com/ericjovian/gin-template/entity"
	mocks "github.com/ericjovian/gin-template/mocks/usecase"
	"github.com/ericjovian/gin-template/server"
	"github.com/ericjovian/gin-template/testutils"
	"github.com/ericjovian/gin-template/utils"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

var (
	book = entity.Book{
		Title:       "Narnia",
		Description: "A book of Narnia",
		Quantity:    10,
		AuthorId:    2,
	}
	reqBook = dto.BookRequest{
		Title:       "Narnia",
		Description: "A book of Narnia",
		Quantity:    10,
		AuthorId:    2,
	}
	resBook = dto.BookResponse{
		ID:          1,
		Title:       "Narnia",
		Description: "A book of Narnia",
		Quantity:    10,
		AuthorId:    2,
	}
)

func TestHandlerGetBooks(t *testing.T) {
	t.Run("should return list of book when status code 200", func(t *testing.T) {
		books := []*entity.Book{
			&book,
		}
		expectedResult := map[string]interface{}{
			"data": books,
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		mockUsecase.On("GetBooks").Return(books, nil)
		cfg := &server.RouterConfig{
			BookUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code INTERNAL_SERVER_ERROR when error occured", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeInternalServerError,
			"message": "error",
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		mockUsecase.On("GetBooks").Return([]*entity.Book{}, errors.New("error"))
		cfg := &server.RouterConfig{
			BookUsecase: mockUsecase,
		}

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		_, rec := testutils.ServeReq(cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})
}

func TestHandlerCreateBook(t *testing.T) {
	t.Run("should return inserted book when status code 201", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"data": resBook,
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		mockUsecase.On("CreateBook", reqBook).Return(&resBook, nil)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(book)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when title is missing from body", func(t *testing.T) {
		invalidBody := entity.Book{
			Description: "A book of Narnia",
			Quantity:    10,
			AuthorId:    2,
		}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidBody)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when description is missing from body", func(t *testing.T) {
		invalidBody := entity.Book{
			Title:    "Narnia",
			Quantity: 10,
			AuthorId: 2,
		}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidBody)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when description less than 10 characters", func(t *testing.T) {
		invalidBody := entity.Book{
			Title:       "Narnia",
			Description: "A book of",
			Quantity:    10,
			AuthorId:    2,
		}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidBody)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when quantity less or equal 0", func(t *testing.T) {
		invalidBody := entity.Book{
			Title:       "Narnia",
			Description: "A book of Narnia",
			Quantity:    0,
			AuthorId:    2,
		}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidBody)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code BAD_REQUEST when author id is missing from body", func(t *testing.T) {
		invalidBody := entity.Book{
			Title:       "Narnia",
			Description: "A book of Narnia",
			Quantity:    10,
		}
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeBadRequest,
			"message": utils.ErrInvalidBody.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(invalidBody)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})

	t.Run("should return error code DUPLICATE_RECORD when trying to insert existing title", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"code":    utils.ErrCodeDuplicate,
			"message": utils.ErrDuplicateBook.Error(),
		}
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		mockUsecase := mocks.NewBookUsecase(t)
		mockUsecase.On("CreateBook", reqBook).Return(nil, utils.ErrDuplicateBook)
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(book)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
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
		mockUsecase := mocks.NewBookUsecase(t)
		mockUsecase.On("CreateBook", reqBook).Return(nil, errors.New("error"))
		cfg := server.RouterConfig{
			BookUsecase: mockUsecase,
		}
		payload := testutils.MakeRequestBody(book)

		req, _ := http.NewRequest(http.MethodPost, "/books", payload)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(jsonExpectedResult), rec.Body.String())
	})
}
