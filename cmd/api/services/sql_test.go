package services

import (
	"rmansilla92/agree-coding-challenge/cmd/api/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *ServicesTestSuite) Test_GetCardFromDB() {
	suite.T().Run("success", func(t *testing.T) {
		id := int64(1)
		name := "Mago Oscuro"
		serieCode := "46986414"
		subtype := "Monstruos Normales"
		typeDescription := "Monstruo"
		price := float64(100)
		dateCreated := "2002-03-08T00:00:00Z"
		expected := &domain.CardEntity{
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
		actual, err := suite.sql.GetCardFromDB("1")
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), expected, actual)
	})
}
