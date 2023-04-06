package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/fydhfzh/assignment-3/dto"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/service"
	"github.com/gin-gonic/gin"
)

type weatherHandler struct {
	weatherService service.WeatherService
}

type WeatherHandler interface {
	CreateWeather(ctx *gin.Context)
	GetLastWeather(ctx *gin.Context)
	ShowWeather(ctx *gin.Context)
	UpdateWeatherEachTime()
}

func NewWeatherHandler(weatherService service.WeatherService) WeatherHandler {
	return &weatherHandler{
		weatherService: weatherService,
	}
}

func (w *weatherHandler) CreateWeather(ctx *gin.Context) {
	
	var weatherPayload dto.WeatherRequest

	if err := ctx.ShouldBindJSON(&weatherPayload); err != nil {
		errBindJSON := errs.NewUnprocessableEntityError("invalid request body")

		ctx.JSON(http.StatusUnprocessableEntity, errBindJSON)

		return
	}

	response, err := w.weatherService.CreateWeather(weatherPayload)

	if err != nil {
		ctx.JSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)

}

func (w *weatherHandler) GetLastWeather(ctx *gin.Context) {
	response, err := w.weatherService.GetLastWeather()

	if err != nil {
		ctx.JSON(err.Status(), err)

		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (w *weatherHandler) ShowWeather(ctx *gin.Context) {
	response, err := w.weatherService.GetLastWeather()

	if err != nil {
		ctx.HTML(err.Status(), "error.html", err)

		return
	}

	ctx.HTML(response.StatusCode, "weather.html", response)
}

func (w *weatherHandler) UpdateWeatherEachTime(){
	for{
		water := rand.Intn(101)
		wind := rand.Intn(101)

		randWeatherData := dto.WeatherRequest{
			Water: water,
			Wind: wind,
		}

		response, err := w.weatherService.CreateWeather(randWeatherData)

		if err != nil {
			log.Fatal(err.Message())
		}

		file, errBytes := json.Marshal(response)

		if errBytes != nil {
			log.Fatal(errBytes)
		}

		errBytes = ioutil.WriteFile("weather.json", file, 0644)

		if errBytes != nil {
			log.Fatal(errBytes)
		}

		fmt.Printf(
		`		{
			"water": %d,
			"wind": %d
		}
		status water : %s
		status wind : %s
		`, 
		response.Water, response.Wind, response.WaterStatus, response.WindStatus)

		time.Sleep(time.Second * 5)
	}
}