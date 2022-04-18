package controllers

import (
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
)

func (suite *ControllersTestSuite) TestDeleteCard() {
	suite.T().Run("success", func(t *testing.T) {
       mockServices := new(mocks.Services)
	   mockServices.On("ProcessDeleteCard", "1").Return(nil).Once()
	   suite.controllers = NewControllers(mockServices)
	   r := gin.Default()
	   r.DELETE("/yu-gi-oh/cards/:card_id", suite.controllers.DeleteCard)
	   apitest.New().
	   Handler(r).
	   Delete("/yu-gi-oh/cards/1").
	   Query("card_id", "1").
	   Expect(suite.T()).
	   Status(http.StatusAccepted).
	   End()
	})
}