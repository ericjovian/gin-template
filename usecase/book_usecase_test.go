package usecase_test

import (
	"errors"
	"testing"

	"github.com/ericjovian/gin-template/dto"
	"github.com/ericjovian/gin-template/entity"
	mocks "github.com/ericjovian/gin-template/mocks/repository"
	"github.com/ericjovian/gin-template/usecase"
	"github.com/ericjovian/gin-template/utils"
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
	resBook = book.ToDTO()
)

func TestUsecaseGetBooks(t *testing.T) {
	t.Run("should return list of books when no error occured", func(t *testing.T) {
		books := []*entity.Book{
			&book,
		}
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("Find").Return(books, nil)

		res, err := uc.GetBooks()

		assert.NoError(t, err)
		assert.Equal(t, books, res)
	})
}

func TestGetBookById(t *testing.T) {
	t.Run("should return book when no error occured", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetById", book.ID).Return(&book, nil)

		res, err := uc.GetBookById(book.ID)

		assert.NoError(t, err)
		assert.Equal(t, &book, res)
	})
}

func TestUsecaseCreateBook(t *testing.T) {
	t.Run("should return inserted book when success to insert", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetByTitle", book.Title).Return(nil, nil)
		mockRepo.On("Insert", book).Return(&book, nil)

		res, err := uc.CreateBook(reqBook)

		assert.NoError(t, err)
		assert.Equal(t, resBook, res)
	})

	t.Run("should return error duplicate when book with same title exists", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetByTitle", book.Title).Return(&book, nil)

		_, err := uc.CreateBook(reqBook)

		assert.ErrorIs(t, utils.ErrDuplicateBook, err)
	})

	t.Run("should return error when error get book by title", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetByTitle", book.Title).Return(nil, errors.New("error"))

		_, err := uc.CreateBook(reqBook)

		assert.Error(t, err)
	})

	t.Run("should return error when error insert to db", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookUsecase(&usecase.BookUConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetByTitle", book.Title).Return(nil, nil)
		mockRepo.On("Insert", book).Return(nil, errors.New("error"))

		_, err := uc.CreateBook(reqBook)

		assert.Error(t, err)
	})
}
