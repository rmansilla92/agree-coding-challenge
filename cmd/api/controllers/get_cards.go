package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *controllers) GetCards(c *gin.Context) {
	response, err := ctrl.services.GetCardsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
	return
}
