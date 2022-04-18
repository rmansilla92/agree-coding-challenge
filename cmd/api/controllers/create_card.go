package controllers

import (
	"encoding/json"
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"
	"rmansilla92/agree-coding-challenge/cmd/api/domain"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *controllers) CreateCard(c *gin.Context) {
	newCard := new(domain.CardDTO)
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(newCard); err != nil {
		err := apierrors.NewBadRequestApiError("Error decoding body")
		c.JSON(err.Status(), err)
		return
	}
	dateCreated := time.Now().Format(time.RFC3339)
	dateCreated = dateCreated[0:10]
	newCard.DateCreated = &(dateCreated)
	if err := ctrl.services.ProcessCreateCard(newCard); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"status": "created", "message": "Card created successfully"})
}
