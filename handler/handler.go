package handler

import "github.com/ericjovian/gin-template/usecase"

type Handler struct {
	bookUsecase         usecase.BookUsecase
	borrowRecordUsecase usecase.BorrowRecordUsecase
}

type Config struct {
	BookUsecase         usecase.BookUsecase
	BorrowRecordUsecase usecase.BorrowRecordUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		bookUsecase:         cfg.BookUsecase,
		borrowRecordUsecase: cfg.BorrowRecordUsecase,
	}
}
