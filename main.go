package main

import (
	"log"
	"losevs/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/weather/now/:city", handlers.GetCityNow)
	router.GET("/weather/:city", handlers.GetCityFuture)
	log.Fatalln(router.Run(":80"))
}
