package services

import (
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"
	"rmansilla92/agree-coding-challenge/cmd/api/domain"
	"strconv"
)

type (
	Services interface {
		GetCardsService() (map[string]interface{}, error)
		GetCardService(cardID string) (*domain.CardDTO, apierrors.ApiError)
		ProcessCreateCard(card *domain.CardDTO) apierrors.ApiError
		ProcessUpdateCard(cardID string, card *domain.CardDTO) apierrors.ApiError
		ProcessDeleteCard(cardID string) apierrors.ApiError
	}
	services struct {
		sqlServices SQLServices
	}
)

func NewServices(sqlService SQLServices) Services {
	return &services{sqlServices: sqlService}
}

func (srv *services) GetCardsService() (map[string]interface{}, error) {
	result, err := srv.sqlServices.GetCardsFromDB()
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("error getting cards", err)
	}
	response := make(map[string]interface{})
	cards := make([]*domain.CardDTO, 0) 
	for _, cardEntity := range result {
		cards = append(cards, cardEntity.CardsEntityToDTO())
	}
	response["cards"] = cards
	response["total"] = len(result)
	return response, nil
}

func (srv *services) GetCardService(cardID string) (*domain.CardDTO, apierrors.ApiError) {
	cardEntity, err := srv.sqlServices.GetCardFromDB(cardID)
	if err != nil {
		return nil, apierrors.NewNotFoundApiError(err.Error())
	}
	card := cardEntity.CardsEntityToDTO()
	return card, nil
}

func (srv *services) ProcessCreateCard(card *domain.CardDTO) apierrors.ApiError {
	cardEntity := card.CardsDTOToEntity()
	
	if err := srv.sqlServices.CreateCardIntoDB(cardEntity); err != nil {
		return apierrors.NewInternalServerApiError("error creating card", err)
	}
	return nil
}

func (srv *services) ProcessUpdateCard(cardID string, card *domain.CardDTO) apierrors.ApiError {
	cardEntity := card.CardsDTOToEntity()
	intCardID, _ := strconv.ParseInt(cardID, 0, 64)
	cardEntity.ID = &intCardID
	if err := srv.sqlServices.UpdateCardIntoDB(cardEntity); err != nil {
		return apierrors.NewInternalServerApiError("error updating card", err)
	}
	return nil
}

func (srv *services) ProcessDeleteCard(cardID string) apierrors.ApiError {
	if err := srv.sqlServices.DeleteCardFromDB(cardID); err != nil {
		return apierrors.NewInternalServerApiError("error deleting card", err)
	}
	return nil
}
