package usecase

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/gin-template/dto"
	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
	"git.garena.com/sea-labs-id/batch-05/gin-template/utils"
)

type BorrowRecordUsecase interface {
	CreateBorrowRecord(dto.BorrowRecordRequest) (*dto.BorrowRecordResponse, error)
	ReturnBook(int) (*dto.BorrowRecordResponse, error)
}

type borrowRecordUsecaseImpl struct {
	userUsecase            UserUsecase
	bookUsecase            BookUsecase
	borrowRecordRepository repository.BorrowRecordRepository
}

type BorrowRecordUConfig struct {
	BorrowRecordRepository repository.BorrowRecordRepository
	UserUsecase            UserUsecase
	BookUsecase            BookUsecase
}

func NewBorrowRecordUsecase(cfg *BorrowRecordUConfig) BorrowRecordUsecase {
	return &borrowRecordUsecaseImpl{
		borrowRecordRepository: cfg.BorrowRecordRepository,
		userUsecase:            cfg.UserUsecase,
		bookUsecase:            cfg.BookUsecase,
	}
}

func (u *borrowRecordUsecaseImpl) CreateBorrowRecord(req dto.BorrowRecordRequest) (*dto.BorrowRecordResponse, error) {
	_, err := u.userUsecase.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}

	bookExist, err := u.bookUsecase.GetBookById(req.BookId)
	if err != nil {
		return nil, err
	}

	if bookExist.Quantity == 0 {
		return nil, utils.ErrEmptyBook
	}

	var borrowRecord entity.BorrowRecord
	borrowRecord.ToStruct(req)
	res, err := u.borrowRecordRepository.Insert(borrowRecord)
	if err != nil {
		return nil, err
	}

	return res.ToDTO(), nil
}

func (u *borrowRecordUsecaseImpl) ReturnBook(id int) (*dto.BorrowRecordResponse, error) {
	record, err := u.borrowRecordRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if record.Status == entity.StatusReturned {
		return nil, utils.ErrAlreadyReturned
	}

	record.Status = entity.StatusReturned
	now := time.Now()
	record.ReturnDate = &now

	res, err := u.borrowRecordRepository.UpdateOnReturn(*record)
	if err != nil {
		return nil, err
	}

	return res.ToDTO(), nil
}
