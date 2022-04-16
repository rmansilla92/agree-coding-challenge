package main

import (
	"fmt"
	"rmansilla92/agree-coding-challenge/cmd/api/config"
	"rmansilla92/agree-coding-challenge/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	controller := controllers.NewControllers()
	if err := run("8080", controller); err != nil {
		fmt.Println("error running server")
	}
}

func run(port string, controller controllers.Controllers) error {
	router := config.CustomRouter()
	mapRoutes(router, controller)
	return router.Run(":" + port)
}

func mapRoutes(router *gin.Engine, controller controllers.Controllers) {
	router.GET("/ping", controller.Ping)
}
