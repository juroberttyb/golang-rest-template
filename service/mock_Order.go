// Code generated by mockery v2.41.0. DO NOT EDIT.

package service

import (
	context "context"

	models "github.com/A-pen-app/kickstart/models"
	mock "github.com/stretchr/testify/mock"
)

// MockOrder is an autogenerated mock type for the Order type
type MockOrder struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, orderID
func (_m *MockOrder) Delete(ctx context.Context, orderID string) error {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orderID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBoard provides a mock function with given fields: ctx, boardType
func (_m *MockOrder) GetBoard(ctx context.Context, boardType models.OrderBoardType) (*models.Board, string, error) {
	ret := _m.Called(ctx, boardType)

	if len(ret) == 0 {
		panic("no return value specified for GetBoard")
	}

	var r0 *models.Board
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderBoardType) (*models.Board, string, error)); ok {
		return rf(ctx, boardType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderBoardType) *models.Board); ok {
		r0 = rf(ctx, boardType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Board)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.OrderBoardType) string); ok {
		r1 = rf(ctx, boardType)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, models.OrderBoardType) error); ok {
		r2 = rf(ctx, boardType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Make provides a mock function with given fields: ctx, action, price, amount
func (_m *MockOrder) Make(ctx context.Context, action models.OrderAction, price int, amount int) error {
	ret := _m.Called(ctx, action, price, amount)

	if len(ret) == 0 {
		panic("no return value specified for Make")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction, int, int) error); ok {
		r0 = rf(ctx, action, price, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Take provides a mock function with given fields: ctx, action, amount
func (_m *MockOrder) Take(ctx context.Context, action models.OrderAction, amount int) error {
	ret := _m.Called(ctx, action, amount)

	if len(ret) == 0 {
		panic("no return value specified for Take")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.OrderAction, int) error); ok {
		r0 = rf(ctx, action, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockOrder creates a new instance of MockOrder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOrder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOrder {
	mock := &MockOrder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
