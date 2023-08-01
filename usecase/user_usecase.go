package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
)

type UserUsecase interface {
	GetUserById(int) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserUConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(cfg *UserUConfig) UserUsecase {
	return &userUsecaseImpl{userRepository: cfg.UserRepository}
}

func (u *userUsecaseImpl) GetUserById(id int) (*entity.User, error) {
	return u.userRepository.GetById(id)
}
