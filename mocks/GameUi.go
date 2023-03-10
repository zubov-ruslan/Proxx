// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GameUi is an autogenerated mock type for the GameUi type
type GameUi struct {
	mock.Mock
}

// DisplayErrorMessage provides a mock function with given fields: msg
func (_m *GameUi) DisplayErrorMessage(msg string) {
	_m.Called(msg)
}

// DisplaySuccessMessage provides a mock function with given fields: msg
func (_m *GameUi) DisplaySuccessMessage(msg string) {
	_m.Called(msg)
}

// GetCoordinations provides a mock function with given fields:
func (_m *GameUi) GetCoordinations() (uint8, uint8, error) {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 uint8
	if rf, ok := ret.Get(1).(func() uint8); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(uint8)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Render provides a mock function with given fields: isHideNotOpened
func (_m *GameUi) Render(isHideNotOpened bool) {
	_m.Called(isHideNotOpened)
}

type mockConstructorTestingTNewGameUi interface {
	mock.TestingT
	Cleanup(func())
}

// NewGameUi creates a new instance of GameUi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGameUi(t mockConstructorTestingTNewGameUi) *GameUi {
	mock := &GameUi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
