package usecase

import (
	"github.com/ericjovian/gin-template/dto"
	"github.com/ericjovian/gin-template/entity"
	"github.com/ericjovian/gin-template/repository"
	"github.com/ericjovian/gin-template/utils"
)

type BookUsecase interface {
	GetBooks() ([]*entity.Book, error)
	GetBookById(int) (*entity.Book, error)
	CreateBook(dto.BookRequest) (*dto.BookResponse, error)
}

type bookUsecaseImpl struct {
	bookRepository repository.BookRepository
}

type BookUConfig struct {
	BookRepository repository.BookRepository
}

func NewBookUsecase(cfg *BookUConfig) BookUsecase {
	return &bookUsecaseImpl{bookRepository: cfg.BookRepository}
}

func (u *bookUsecaseImpl) GetBooks() ([]*entity.Book, error) {
	return u.bookRepository.Find()
}

func (u *bookUsecaseImpl) GetBookById(id int) (*entity.Book, error) {
	return u.bookRepository.GetById(id)
}

func (u *bookUsecaseImpl) CreateBook(req dto.BookRequest) (*dto.BookResponse, error) {
	exist, err := u.bookRepository.GetByTitle(req.Title)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, utils.ErrDuplicateBook
	}

	var bookToInsert entity.Book
	bookToInsert.ToStruct(req)

	res, err := u.bookRepository.Insert(bookToInsert)
	if err != nil {
		return nil, err
	}

	return res.ToDTO(), nil
}
