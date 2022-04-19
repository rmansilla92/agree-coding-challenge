package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Paths Information 
// GetCards godoc
// @Summary Obtiene la información de una carta especifica de la base de datos
// @ID GetCards
// @Consume application/json
// @Produce json
// @Param card_id formData string true "Card ID"
// @Success 200 {object} domain.CardDTO
// @Failure 404 {object} apierrors.ApiError
// @Router /cards/card_id [get]
func (ctrl *controllers) GetCards(c *gin.Context) {
	response, err := ctrl.services.GetCardsService()
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}
