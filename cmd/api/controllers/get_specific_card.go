package controllers

import (
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"

	"github.com/gin-gonic/gin"
)

func (ctrl *controllers) GetSpecificCard(c *gin.Context) {
	cardID, ok := c.Params.Get("card_id")
	if !ok {
		err := apierrors.NewBadRequestApiError("missing param card_id")
		c.JSON(err.Status(), err)
		return
	}
	response, err := ctrl.services.GetCardService(cardID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, response)
}
