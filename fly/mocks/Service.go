// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Execute provides a mock function with given fields: height
func (_m *Service) Execute(height int) error {
	ret := _m.Called(height)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(height)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
