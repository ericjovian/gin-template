package repository

import (
	"errors"

	"github.com/ericjovian/gin-template/entity"
	"github.com/ericjovian/gin-template/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(int) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(cfg *UserRConfig) UserRepository {
	return &userRepositoryImpl{db: cfg.DB}
}

func (r *userRepositoryImpl) GetById(id int) (*entity.User, error) {
	var res *entity.User
	err := r.db.First(&res, id)
	if err.Error != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrUserNotFound
		}
		return nil, err.Error
	}

	return res, nil
}
