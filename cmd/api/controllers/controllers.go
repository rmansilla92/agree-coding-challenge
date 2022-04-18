package controllers

import (
	"rmansilla92/agree-coding-challenge/cmd/api/services"

	"github.com/gin-gonic/gin"
)

type (
	Controllers interface {
		Ping(c *gin.Context)
		CreateCard(c *gin.Context)
		UpdateCard(c *gin.Context)
		GetCards(c *gin.Context)
		GetSpecificCard(c *gin.Context)
		DeleteCard(c *gin.Context)
	}
	controllers struct {
		services services.Services
	}
)

func NewControllers(service services.Services) Controllers {
	return &controllers{services: service}
}
