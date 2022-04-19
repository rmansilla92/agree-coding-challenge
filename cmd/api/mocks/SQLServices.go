// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	domain "rmansilla92/agree-coding-challenge/cmd/api/domain"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// SQLServices is an autogenerated mock type for the SQLServices type
type SQLServices struct {
	mock.Mock
}

// CreateCardIntoDB provides a mock function with given fields: cardEntity
func (_m *SQLServices) CreateCardIntoDB(cardEntity *domain.NewCardEntity) error {
	ret := _m.Called(cardEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.NewCardEntity) error); ok {
		r0 = rf(cardEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCardFromDB provides a mock function with given fields: cardID
func (_m *SQLServices) DeleteCardFromDB(cardID string) error {
	ret := _m.Called(cardID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(cardID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCardFromDB provides a mock function with given fields: cardID
func (_m *SQLServices) GetCardFromDB(cardID string) (*domain.CardEntity, error) {
	ret := _m.Called(cardID)

	var r0 *domain.CardEntity
	if rf, ok := ret.Get(0).(func(string) *domain.CardEntity); ok {
		r0 = rf(cardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CardEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCardsFromDB provides a mock function with given fields:
func (_m *SQLServices) GetCardsFromDB() ([]domain.CardEntity, error) {
	ret := _m.Called()

	var r0 []domain.CardEntity
	if rf, ok := ret.Get(0).(func() []domain.CardEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.CardEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCardIntoDB provides a mock function with given fields: cardEntity
func (_m *SQLServices) UpdateCardIntoDB(cardEntity *domain.NewCardEntity) error {
	ret := _m.Called(cardEntity)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.NewCardEntity) error); ok {
		r0 = rf(cardEntity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSQLServices creates a new instance of SQLServices. It also registers a cleanup function to assert the mocks expectations.
func NewSQLServices(t testing.TB) *SQLServices {
	mock := &SQLServices{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}