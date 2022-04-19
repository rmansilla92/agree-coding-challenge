package main

import (
	"fmt"
	"os"
	"rmansilla92/agree-coding-challenge/cmd/api/config"
	"rmansilla92/agree-coding-challenge/cmd/api/controllers"
	"rmansilla92/agree-coding-challenge/cmd/api/docs"
	"rmansilla92/agree-coding-challenge/cmd/api/services"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Agree – Backend Engineer Coding Challenge"
	docs.SwaggerInfo.Description = "Esta API tiene los endpoints relacionados con las cartas Yu-Gi-Oh! solicitados en el challenge"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://agree-coding-challenge-2-eissifkqta-rj.a.run.app"
	docs.SwaggerInfo.BasePath = "/yu-gi-oh"
	docs.SwaggerInfo.Schemes = []string{"https"}
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
