package handler

import (
	"github.com/fydhfzh/assignment-3/database"
	"github.com/fydhfzh/assignment-3/repository/weather_repository/weather_pg"
	"github.com/fydhfzh/assignment-3/service"
	"github.com/gin-gonic/gin"
)

const PORT = ":3000"

func StartApp() {
	database.InitializeDB()
	
	db := database.GetInstance()

	weatherRepo := weather_pg.NewWeatherPG(db)
	weatherService := service.NewWeatherService(weatherRepo)
	weatherHandler := NewWeatherHandler(weatherService)

	
	route := gin.Default()

	go weatherHandler.UpdateWeatherEachTime()

	route.LoadHTMLGlob("assets/html/*.html")

	route.POST("/weather", weatherHandler.CreateWeather)
	route.GET("/api/weather", weatherHandler.GetLastWeather)
	route.GET("/weather", weatherHandler.ShowWeather)
	route.Run(PORT)
}