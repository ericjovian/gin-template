package server

import (
	"github.com/ericjovian/gin-template/handler"
	"github.com/ericjovian/gin-template/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookUsecase         usecase.BookUsecase
	BorrowRecordUsecase usecase.BorrowRecordUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		BookUsecase:         cfg.BookUsecase,
		BorrowRecordUsecase: cfg.BorrowRecordUsecase,
	})

	router.GET("/books", h.GetBooks)
	router.POST("/books", h.CreateBook)

	router.POST("/borrow-records", h.CreateBorrowRecord)
	router.PUT("/borrow-records/:id", h.ReturnBorrowRecord)

	return router
}
