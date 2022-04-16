package controllers

import "github.com/gin-gonic/gin"

type (
	Controllers interface {
		Ping(c *gin.Context)
	}
	controllers struct {
	}
)

func NewControllers() Controllers {
	return &controllers{}
}
