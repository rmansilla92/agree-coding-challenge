package services

import (
	"rmansilla92/agree-coding-challenge/cmd/api/domain"
	"rmansilla92/agree-coding-challenge/cmd/api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *ServicesTestSuite) Test_GetCardService() {
	suite.T().Run("success", func (t *testing.T) {
		mockSQLService := new(mocks.SQLServices)
		id := int64(1)
		name := "Mago Oscuro"
		serieCode := "46986414"
		subtype := "Monstruos Normales"
		typeDescription := "Monstruo"
		price := float64(100)
		dateCreated := "2002-03-08T00:00:00Z"
		expectedCardEntity := &domain.CardEntity{
			ID:                 &id,
			Name:               &name,
			FirstEdition:       true,
			SerieCode:          &serieCode,
			SubTypeID:          &id,
			SubtypeDescription: &subtype,
			TypeID:             &id,
			TypeDescription:    &typeDescription,
			ATK:                2500,
			DEF:                2100,
			Stars:              7,
			Description:        "El más grande de los magos en cuanto al ataque y la defensa",
			Price:              &price,
			ImageID:            &id,
			ImageDescription:   &name,
			DateCreated:        &dateCreated,
		}
		mockSQLService.On("GetCardFromDB", "1").Return(expectedCardEntity, nil).Once()
		expectedCardDTO := expectedCardEntity.CardsEntityToDTO()
		suite.services = NewServices(mockSQLService)
		actual, err := suite.services.GetCardService("1")
		assert.Equal(suite.T(), expectedCardDTO, actual)
		assert.Nil(suite. T(), err)
	})
}