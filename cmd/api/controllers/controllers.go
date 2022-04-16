package controllers

import "github.com/gin-gonic/gin"

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
	}
)

func NewControllers() Controllers {
	return &controllers{}
}
