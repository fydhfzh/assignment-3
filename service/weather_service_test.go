package service

import (
	"net/http"
	"testing"

	"github.com/fydhfzh/assignment-3/dto"
	"github.com/fydhfzh/assignment-3/entity"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/repository/weather_repository"
	"github.com/stretchr/testify/assert"
)

func TestWeatherService_CreateNewWeather_Success(t *testing.T){
	payload := dto.WeatherRequest{
		Water: 18,
		Wind: 20,
	}

	weatherRepo := weather_repository.NewWeatherMockRepository()

	weatherService := NewWeatherService(weatherRepo)

	weather_repository.CreateWeather = func(weather entity.Weather) (*entity.Weather, errs.ErrMessage) {
		weatherData := entity.Weather{
			Water: 18,
			Wind: 18,
			WaterStatus: "Bahaya",
			WindStatus: "Bahaya",
		}

		return &weatherData, nil
	}

	response, err := weatherService.CreateWeather(payload)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 18, response.Water)
	assert.Equal(t, 18, response.Wind)
	assert.Equal(t, "Bahaya", response.WaterStatus)
	assert.Equal(t, "Bahaya", response.WindStatus)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestWeatherService_CreateNewWeather_BadRequestError(t *testing.T) {
	payload := dto.WeatherRequest{
		Water: 18,
		Wind: 20,
	}

	weatherRepo := weather_repository.NewWeatherMockRepository()

	weatherService := NewWeatherService(weatherRepo)

	weather_repository.CreateWeather = func(weather entity.Weather) (*entity.Weather, errs.ErrMessage) {
		return nil, errs.NewBadRequest("bad request body")
	}

	response, err := weatherService.CreateWeather(payload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, "bad request body", err.Message())
	assert.Equal(t, "BAD_REQUEST", err.Error())
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestWeatherService_CreateNewWeather_DbError(t *testing.T){
	payload := dto.WeatherRequest{
		Water: 18,
		Wind: 20,
	}

	weatherRepo := weather_repository.NewWeatherMockRepository()

	weatherService := NewWeatherService(weatherRepo)

	weather_repository.CreateWeather = func(weather entity.Weather) (*entity.Weather, errs.ErrMessage) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := weatherService.CreateWeather(payload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, "something went wrong", err.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestWeatherService_GetLastWeather_Success(t *testing.T) {

	weatherRepo := weather_repository.NewWeatherMockRepository()

	weatherService := NewWeatherService(weatherRepo)

	weather_repository.GetLastWeather = func() (*entity.Weather, errs.ErrMessage) {
		weatherData := entity.Weather{
			Water: 18,
			Wind: 18,
			WaterStatus: "Bahaya",
			WindStatus: "Bahaya",
		}

		return &weatherData, nil
	}

	response, err := weatherService.GetLastWeather()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 18, response.Water)
	assert.Equal(t, 18, response.Wind)
	assert.Equal(t, "Bahaya", response.WaterStatus)
	assert.Equal(t, "Bahaya", response.WindStatus)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestWeatherService_GetLastWeather_DbError(t *testing.T) {
	weatherRepo := weather_repository.NewWeatherMockRepository()

	weatherService := NewWeatherService(weatherRepo)

	weather_repository.GetLastWeather = func() (*entity.Weather, errs.ErrMessage) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := weatherService.GetLastWeather()

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, "something went wrong", err.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}