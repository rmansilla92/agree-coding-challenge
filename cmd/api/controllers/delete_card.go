package controllers

import (
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"

	"github.com/gin-gonic/gin"
)

func (ctrl *controllers) DeleteCard(c *gin.Context) {
	cardID, ok := c.Params.Get("card_id")
	if !ok {
		err := apierrors.NewBadRequestApiError("missing param card_id")
		c.JSON(err.Status(), err)
		return
	}
	if err := ctrl.services.ProcessDeleteCard(cardID); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusAccepted, map[string]string{"status": "accepted", "message": "Card deleted successfully"})
}
