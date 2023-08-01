package entity

import (
	"time"

	"github.com/ericjovian/gin-template/dto"
	"gorm.io/gorm"
)

type BorrowRecord struct {
	ID         int
	UserId     int
	BookId     int
	Status     string
	BorrowDate time.Time
	ReturnDate *time.Time
	gorm.Model
}

const (
	StatusBorrowed = "borrow"
	StatusReturned = "return"
)

func (br *BorrowRecord) ToDTO() *dto.BorrowRecordResponse {
	return &dto.BorrowRecordResponse{
		ID:         br.ID,
		UserId:     br.UserId,
		BookId:     br.BookId,
		Status:     br.Status,
		BorrowDate: br.BorrowDate,
		ReturnDate: br.ReturnDate,
	}
}

func (br *BorrowRecord) ToStruct(req dto.BorrowRecordRequest) {
	br.UserId = req.UserId
	br.BookId = req.BookId
	br.Status = StatusBorrowed
	br.BorrowDate = time.Now()
}
