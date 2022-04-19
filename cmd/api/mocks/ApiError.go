// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	apierrors "rmansilla92/agree-coding-challenge/cmd/api/apierrors"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ApiError is an autogenerated mock type for the ApiError type
type ApiError struct {
	mock.Mock
}

// Cause provides a mock function with given fields:
func (_m *ApiError) Cause() apierrors.CauseList {
	ret := _m.Called()

	var r0 apierrors.CauseList
	if rf, ok := ret.Get(0).(func() apierrors.CauseList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.CauseList)
		}
	}

	return r0
}

// Code provides a mock function with given fields:
func (_m *ApiError) Code() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *ApiError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *ApiError) Message() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *ApiError) Status() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewApiError creates a new instance of ApiError. It also registers a cleanup function to assert the mocks expectations.
func NewApiError(t testing.TB) *ApiError {
	mock := &ApiError{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}