package server

import (
	"log"

	"github.com/ericjovian/gin-template/db"
	"github.com/ericjovian/gin-template/repository"
	"github.com/ericjovian/gin-template/usecase"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	bookRepo := repository.NewBookRepository(&repository.BookRConfig{
		DB: db.Get(),
	})
	bookUsecase := usecase.NewBookUsecase(&usecase.BookUConfig{
		BookRepository: bookRepo,
	})

	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(),
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepository: userRepo,
	})

	borrowRepo := repository.NewBorrowRecordRepository(&repository.BorrowRecordRConfig{
		DB:       db.Get(),
		BookRepo: bookRepo,
	})
	borrowRecordUsecase := usecase.NewBorrowRecordUsecase(&usecase.BorrowRecordUConfig{
		BorrowRecordRepository: borrowRepo,
		UserUsecase:            userUsecase,
		BookUsecase:            bookUsecase,
	})

	return NewRouter(&RouterConfig{
		BookUsecase:         bookUsecase,
		BorrowRecordUsecase: borrowRecordUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
