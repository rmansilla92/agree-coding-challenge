package mocks

import (
	apierrors "rmansilla92/agree-coding-challenge/cmd/api/apierrors"
	"rmansilla92/agree-coding-challenge/cmd/api/domain"

	mock "github.com/stretchr/testify/mock"
)

type Services struct {
	mock.Mock
}

func (_m *Services) GetCardsService() (map[string]interface{}, error) {
	panic("falta implementar")
}

func (_m *Services) GetCardService(cardID string) (*domain.CardDTO, apierrors.ApiError) {
	panic("falta implementar")
}

func (_m *Services) ProcessCreateCard(card *domain.CardDTO) apierrors.ApiError {
	panic("falta implementar")
}

func (_m *Services) ProcessUpdateCard(cardID string, card *domain.CardDTO) apierrors.ApiError {
	panic("falta implementar")
}

func (_m *Services) ProcessDeleteCard(cardID string) apierrors.ApiError {
	ret := _m.Called(cardID)

	var r0 apierrors.ApiError
	if rf, ok := ret.Get(0).(func(string) apierrors.ApiError); ok {
		r0 = rf(cardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apierrors.ApiError)
		}
	}
	return r0
}
