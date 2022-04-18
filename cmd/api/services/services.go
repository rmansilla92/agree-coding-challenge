package services

import "rmansilla92/agree-coding-challenge/cmd/api/domain"

type (
	Services interface {
		GetCardsService() (map[string]interface{}, error)
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
		return nil, err
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
