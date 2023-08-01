package repository

import (
	"errors"

	"github.com/ericjovian/gin-template/entity"
	"github.com/ericjovian/gin-template/utils"
	"gorm.io/gorm"
)

type BookRepository interface {
	Find() ([]*entity.Book, error)
	GetByTitle(string) (*entity.Book, error)
	GetById(int) (*entity.Book, error)
	Insert(entity.Book) (*entity.Book, error)
	SubstractQuantity(tx *gorm.DB, id, qty int) error
	AddQuantity(tx *gorm.DB, id, qty int) error
}

type bookRepositoryImpl struct {
	db *gorm.DB
}

type BookRConfig struct {
	DB *gorm.DB
}

func NewBookRepository(cfg *BookRConfig) BookRepository {
	return &bookRepositoryImpl{db: cfg.DB}
}

func (r *bookRepositoryImpl) Find() ([]*entity.Book, error) {
	var res []*entity.Book
	err := r.db.Joins("Author").Find(&res)
	if err.Error != nil {
		return nil, err.Error
	}

	return res, nil
}

func (r *bookRepositoryImpl) GetById(id int) (*entity.Book, error) {
	var res *entity.Book
	err := r.db.First(&res, id)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrBookNotFound
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *bookRepositoryImpl) GetByTitle(title string) (*entity.Book, error) {
	var res *entity.Book
	err := r.db.Where("title = ?", title).First(&res)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *bookRepositoryImpl) Insert(req entity.Book) (*entity.Book, error) {
	err := r.db.Create(&req)
	if err.Error != nil {
		return nil, err.Error
	}

	return &req, nil
}

func (r *bookRepositoryImpl) SubstractQuantity(tx *gorm.DB, id, qty int) error {
	err := tx.Model(&entity.Book{}).Where("id = ?", id).Where("quantity > ?", 0).Update("quantity", gorm.Expr("quantity - ?", qty))
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return utils.ErrEmptyBook
	}

	return nil
}

func (r *bookRepositoryImpl) AddQuantity(tx *gorm.DB, id, qty int) error {
	err := tx.Model(&entity.Book{}).Where("id = ?", id).Update("quantity", gorm.Expr("quantity + ?", qty))
	if err.Error != nil {
		return err.Error
	}

	return nil
}
