package main

import (
	"fmt"
	"os"
	"rmansilla92/agree-coding-challenge/cmd/api/config"
	"rmansilla92/agree-coding-challenge/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	controller := controllers.NewControllers()
	if err := run(port, controller); err != nil {
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
	router.POST("/yu-gi-oh/cards", controller.CreateCard)
	router.PUT("/yu-gi-oh/cards/:card_id", controller.UpdateCard)
	router.GET("/yu-gi-oh/cards", controller.GetCards)
	router.GET("/yu-gi-oh/cards/:card_id", controller.GetSpecificCard)
	router.DELETE("/yu-gi-oh/cards/:card_id", controller.DeleteCard)
}
