package dto

import (
	"time"
)

type BorrowRecordRequest struct {
	UserId int `json:"userId" binding:"required"`
	BookId int `json:"bookId" binding:"required"`
}

type BorrowRecordResponse struct {
	ID         int        `json:"id"`
	UserId     int        `json:"userId"`
	BookId     int        `json:"bookId"`
	Status     string     `json:"status"`
	BorrowDate time.Time  `json:"borrowDate"`
	ReturnDate *time.Time `json:"returnDate,omitempty"`
}
