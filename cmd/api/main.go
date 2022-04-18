package main

import (
	"fmt"
	"os"
	"rmansilla92/agree-coding-challenge/cmd/api/config"
	"rmansilla92/agree-coding-challenge/cmd/api/controllers"
	"rmansilla92/agree-coding-challenge/cmd/api/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	if strings.EqualFold(os.Getenv("SCOPE"), "prod") {
		config.LoadProductionValues()
	} else {
		config.LoadDevelopValues()
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db, err := services.OpenDB("default")
	if err != nil {
		panic("Error on init default DB")
	}
	sqlServices := services.NewSQLServices(db)
	services := services.NewServices(sqlServices)
	controllers := controllers.NewControllers(services)
	if err := run("8080", controllers); err != nil {
		fmt.Println("error running server")
	}
}

func run(port string, controllers controllers.Controllers) error {
	router := config.CustomRouter()
	mapRoutes(router, controllers)
	return router.Run(":" + port)
}

func mapRoutes(router *gin.Engine, controllers controllers.Controllers) {
	router.GET("/ping", controllers.Ping)
	router.POST("/yu-gi-oh/cards", controllers.CreateCard)
	router.PUT("/yu-gi-oh/cards/:card_id", controllers.UpdateCard)
	router.GET("/yu-gi-oh/cards", controllers.GetCards)
	router.GET("/yu-gi-oh/cards/:card_id", controllers.GetSpecificCard)
	router.DELETE("/yu-gi-oh/cards/:card_id", controllers.DeleteCard)
}
