// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/ericjovian/gin-template/entity"
	mock "github.com/stretchr/testify/mock"
)

// BorrowRecordRepository is an autogenerated mock type for the BorrowRecordRepository type
type BorrowRecordRepository struct {
	mock.Mock
}

// GetById provides a mock function with given fields: _a0
func (_m *BorrowRecordRepository) GetById(_a0 int) (*entity.BorrowRecord, error) {
	ret := _m.Called(_a0)

	var r0 *entity.BorrowRecord
	if rf, ok := ret.Get(0).(func(int) *entity.BorrowRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowRecord)
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

// Insert provides a mock function with given fields: _a0
func (_m *BorrowRecordRepository) Insert(_a0 entity.BorrowRecord) (*entity.BorrowRecord, error) {
	ret := _m.Called(_a0)

	var r0 *entity.BorrowRecord
	if rf, ok := ret.Get(0).(func(entity.BorrowRecord) *entity.BorrowRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.BorrowRecord) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOnReturn provides a mock function with given fields: _a0
func (_m *BorrowRecordRepository) UpdateOnReturn(_a0 entity.BorrowRecord) (*entity.BorrowRecord, error) {
	ret := _m.Called(_a0)

	var r0 *entity.BorrowRecord
	if rf, ok := ret.Get(0).(func(entity.BorrowRecord) *entity.BorrowRecord); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.BorrowRecord) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowRecordRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowRecordRepository creates a new instance of BorrowRecordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowRecordRepository(t mockConstructorTestingTNewBorrowRecordRepository) *BorrowRecordRepository {
	mock := &BorrowRecordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
