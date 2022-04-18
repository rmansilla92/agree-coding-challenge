package controllers

import (
	"encoding/json"
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"
	"rmansilla92/agree-coding-challenge/cmd/api/domain"

	"github.com/gin-gonic/gin"
)

func (ctrl *controllers) UpdateCard(c *gin.Context) {
	cardID, ok := c.Params.Get("card_id")
	if !ok {
		err := apierrors.NewBadRequestApiError("missing param card_id")
		c.JSON(err.Status(), err)
		return
	}
	updatedCard := new(domain.CardDTO)
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(updatedCard); err != nil {
		err := apierrors.NewBadRequestApiError("Error decoding body")
		c.JSON(err.Status(), err)
		return
	}
	if err := ctrl.services.ProcessUpdateCard(cardID, updatedCard); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusAccepted, map[string]string{"status": "accepted", "message": "Card updated successfully"})
}
