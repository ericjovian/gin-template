package repository

import (
	"errors"

	"github.com/ericjovian/gin-template/entity"
	"github.com/ericjovian/gin-template/utils"
	"gorm.io/gorm"
)

type BorrowRecordRepository interface {
	Insert(entity.BorrowRecord) (*entity.BorrowRecord, error)
	UpdateOnReturn(entity.BorrowRecord) (*entity.BorrowRecord, error)
	GetById(int) (*entity.BorrowRecord, error)
}

type borrowRecordRepositoryImpl struct {
	db       *gorm.DB
	bookRepo BookRepository
}

type BorrowRecordRConfig struct {
	DB       *gorm.DB
	BookRepo BookRepository
}

func NewBorrowRecordRepository(cfg *BorrowRecordRConfig) BorrowRecordRepository {
	return &borrowRecordRepositoryImpl{
		db:       cfg.DB,
		bookRepo: cfg.BookRepo,
	}
}

func (r *borrowRecordRepositoryImpl) GetById(id int) (*entity.BorrowRecord, error) {
	var res *entity.BorrowRecord
	err := r.db.First(&res, id)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrRecordNotFound
		}
		return nil, err.Error
	}

	return res, nil
}

func (r *borrowRecordRepositoryImpl) Insert(req entity.BorrowRecord) (*entity.BorrowRecord, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := tx.Create(&req)
	if err.Error != nil {
		tx.Rollback()
		return nil, err.Error
	}

	qtyToSub := 1
	errSub := r.bookRepo.SubstractQuantity(tx, req.BookId, qtyToSub)
	if errSub != nil {
		tx.Rollback()
		return nil, errSub
	}

	tx.Commit()
	return &req, nil
}

func (r *borrowRecordRepositoryImpl) UpdateOnReturn(req entity.BorrowRecord) (*entity.BorrowRecord, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	err := tx.Save(&req)
	if err.Error != nil {
		tx.Rollback()
		return nil, err.Error
	}

	qtyToAdd := 1
	errSub := r.bookRepo.AddQuantity(tx, req.BookId, qtyToAdd)
	if errSub != nil {
		tx.Rollback()
		return nil, errSub
	}

	tx.Commit()
	return &req, nil
}
