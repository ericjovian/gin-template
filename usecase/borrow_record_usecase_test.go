package usecase_test

import (
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/ericjovian/gin-template/dto"
	"github.com/ericjovian/gin-template/entity"
	mocks "github.com/ericjovian/gin-template/mocks/repository"
	mockUsecase "github.com/ericjovian/gin-template/mocks/usecase"
	"github.com/ericjovian/gin-template/usecase"
	"github.com/ericjovian/gin-template/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateBorrowRecord(t *testing.T) {
	user := entity.User{
		ID:    1,
		Name:  "test",
		Email: "test",
		Phone: "081234567890",
	}
	book := entity.Book{
		ID:          1,
		Title:       "test",
		Description: "test",
		Quantity:    10,
		Cover:       nil,
		AuthorId:    1,
	}
	reqBorrowRecord := dto.BorrowRecordRequest{
		UserId: 1,
		BookId: 1,
	}
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	now := time.Now()
	borrowRecord := entity.BorrowRecord{
		UserId:     1,
		BookId:     1,
		Status:     entity.StatusBorrowed,
		BorrowDate: now,
	}
	resBorrowRecord := borrowRecord.ToDTO()

	t.Run("should return borrow record when no error occured", func(t *testing.T) {
		mockUserUsecase := mockUsecase.NewUserUsecase(t)
		mockBookUsecase := mockUsecase.NewBookUsecase(t)
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockUserUsecase.On("GetUserById", user.ID).Return(&user, nil)
		mockBookUsecase.On("GetBookById", book.ID).Return(&book, nil)
		mockRepo.On("Insert", borrowRecord).Return(&borrowRecord, nil)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
			BookUsecase:            mockBookUsecase,
			UserUsecase:            mockUserUsecase,
		})

		res, err := uc.CreateBorrowRecord(reqBorrowRecord)

		assert.NoError(t, err)
		assert.Equal(t, resBorrowRecord, res)
	})

	t.Run("should return error when user doesn't exist", func(t *testing.T) {
		mockUserUsecase := mockUsecase.NewUserUsecase(t)
		mockBookUsecase := mockUsecase.NewBookUsecase(t)
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockUserUsecase.On("GetUserById", user.ID).Return(nil, utils.ErrUserNotFound)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
			BookUsecase:            mockBookUsecase,
			UserUsecase:            mockUserUsecase,
		})

		_, err := uc.CreateBorrowRecord(reqBorrowRecord)

		assert.ErrorIs(t, utils.ErrUserNotFound, err)
	})

	t.Run("should return error when book doesn't exist", func(t *testing.T) {
		mockUserUsecase := mockUsecase.NewUserUsecase(t)
		mockBookUsecase := mockUsecase.NewBookUsecase(t)
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockUserUsecase.On("GetUserById", user.ID).Return(&user, nil)
		mockBookUsecase.On("GetBookById", book.ID).Return(nil, utils.ErrBookNotFound)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
			BookUsecase:            mockBookUsecase,
			UserUsecase:            mockUserUsecase,
		})

		_, err := uc.CreateBorrowRecord(reqBorrowRecord)

		assert.ErrorIs(t, utils.ErrBookNotFound, err)
	})

	t.Run("should return error when book's stock is empty", func(t *testing.T) {
		emptyBook := entity.Book{
			Quantity: 0,
		}
		mockUserUsecase := mockUsecase.NewUserUsecase(t)
		mockBookUsecase := mockUsecase.NewBookUsecase(t)
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockUserUsecase.On("GetUserById", user.ID).Return(&user, nil)
		mockBookUsecase.On("GetBookById", book.ID).Return(&emptyBook, nil)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
			BookUsecase:            mockBookUsecase,
			UserUsecase:            mockUserUsecase,
		})

		_, err := uc.CreateBorrowRecord(reqBorrowRecord)

		assert.ErrorIs(t, utils.ErrEmptyBook, err)
	})

	t.Run("should return error when error occured on insert", func(t *testing.T) {
		expectedError := errors.New("error")
		mockUserUsecase := mockUsecase.NewUserUsecase(t)
		mockBookUsecase := mockUsecase.NewBookUsecase(t)
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockUserUsecase.On("GetUserById", user.ID).Return(&user, nil)
		mockBookUsecase.On("GetBookById", book.ID).Return(&book, nil)
		mockRepo.On("Insert", borrowRecord).Return(nil, expectedError)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
			BookUsecase:            mockBookUsecase,
			UserUsecase:            mockUserUsecase,
		})

		_, err := uc.CreateBorrowRecord(reqBorrowRecord)

		assert.ErrorIs(t, expectedError, err)
	})
}

func TestReturnBook(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	})
	now := time.Now()

	t.Run("should return updated borrow record when no error occured", func(t *testing.T) {
		reqBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusBorrowed,
			BorrowDate: now,
		}
		updatedBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusReturned,
			BorrowDate: now,
			ReturnDate: &now,
		}
		resBorrowRecord := updatedBorrowRecord.ToDTO()
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockRepo.On("GetById", reqBorrowRecord.ID).Return(&reqBorrowRecord, nil)
		mockRepo.On("UpdateOnReturn", updatedBorrowRecord).Return(&updatedBorrowRecord, nil)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
		})

		res, err := uc.ReturnBook(reqBorrowRecord.ID)

		assert.NoError(t, err)
		assert.Equal(t, resBorrowRecord, res)
	})

	t.Run("should return error when record not found", func(t *testing.T) {
		reqBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusBorrowed,
			BorrowDate: now,
		}
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockRepo.On("GetById", reqBorrowRecord.ID).Return(nil, utils.ErrRecordNotFound)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
		})

		_, err := uc.ReturnBook(reqBorrowRecord.ID)

		assert.ErrorIs(t, utils.ErrRecordNotFound, err)
	})

	t.Run("should return error when borrow status is already returned", func(t *testing.T) {
		reqBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusBorrowed,
			BorrowDate: now,
		}
		returnedBorrowRecord := entity.BorrowRecord{
			Status: entity.StatusReturned,
		}
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockRepo.On("GetById", reqBorrowRecord.ID).Return(&returnedBorrowRecord, nil)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
		})

		_, err := uc.ReturnBook(reqBorrowRecord.ID)

		assert.ErrorIs(t, utils.ErrAlreadyReturned, err)
	})

	t.Run("should return error when error occured on UpdateOnReturn", func(t *testing.T) {
		reqBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusBorrowed,
			BorrowDate: now,
		}
		updatedBorrowRecord := entity.BorrowRecord{
			ID:         1,
			UserId:     1,
			BookId:     1,
			Status:     entity.StatusReturned,
			BorrowDate: now,
			ReturnDate: &now,
		}
		expectedError := errors.New("error")
		mockRepo := mocks.NewBorrowRecordRepository(t)
		mockRepo.On("GetById", reqBorrowRecord.ID).Return(&reqBorrowRecord, nil)
		mockRepo.On("UpdateOnReturn", updatedBorrowRecord).Return(nil, expectedError)
		uc := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
			BorrowRecordRepository: mockRepo,
		})

		_, err := uc.ReturnBook(reqBorrowRecord.ID)

		assert.ErrorIs(t, expectedError, err)
	})
}
