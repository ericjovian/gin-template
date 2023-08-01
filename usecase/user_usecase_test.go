package usecase_test

import (
	"testing"

	"github.com/ericjovian/gin-template/entity"
	mocks "github.com/ericjovian/gin-template/mocks/repository"
	"github.com/ericjovian/gin-template/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	t.Run("should return book when no error occured", func(t *testing.T) {
		user := entity.User{
			ID:    1,
			Name:  "test",
			Email: "test",
			Phone: "081234567890",
		}
		mockRepo := mocks.NewUserRepository(t)
		uc := usecase.NewUserUsecase(&usecase.UserUConfig{
			UserRepository: mockRepo,
		})
		mockRepo.On("GetById", user.ID).Return(&user, nil)

		res, err := uc.GetUserById(user.ID)

		assert.NoError(t, err)
		assert.Equal(t, &user, res)
	})
}
