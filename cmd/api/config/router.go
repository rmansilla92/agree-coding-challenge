package config

import (
	"fmt"
	"net/http"
	"rmansilla92/agree-coding-challenge/cmd/api/apierrors"

	"github.com/gin-gonic/gin"
)

func CustomRouter() *gin.Engine {
	router := gin.New()
	router.NoRoute(noRouteHandler)
	return router
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, apierrors.NewNotFoundApiError(fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)))
}
