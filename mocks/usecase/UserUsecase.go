// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/ericjovian/gin-template/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: _a0
func (_m *UserUsecase) GetUserById(_a0 int) (*entity.User, error) {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
