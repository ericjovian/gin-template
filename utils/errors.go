package utils

import "errors"

const (
	ErrSqlUniqueViolation      = "23505"
	ErrCodeDuplicate           = "DUPLICATE_RECORD"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest          = "BAD_REQUEST"
)

var (
	ErrDuplicateBook   = errors.New("duplicate insertion of book with same title")
	ErrInvalidBody     = errors.New("invalid body request")
	ErrInvalidParam    = errors.New("invalid params")
	ErrUserNotFound    = errors.New("user not found")
	ErrBookNotFound    = errors.New("book not found")
	ErrEmptyBook       = errors.New("book's stock is empty")
	ErrRecordNotFound  = errors.New("record not found")
	ErrAlreadyReturned = errors.New("book already returned")
)
