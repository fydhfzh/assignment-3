package service

import (
	"net/http"

	"github.com/fydhfzh/assignment-3/dto"
	"github.com/fydhfzh/assignment-3/entity"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/pkg/helpers"
	"github.com/fydhfzh/assignment-3/repository/weather_repository"
)

type weatherService struct {
	weatherRepo weather_repository.WeatherRepo
}

type WeatherService interface {
	CreateWeather(weatherPayload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage)
	GetLastWeather() (*dto.WeatherResponse, errs.ErrMessage)
}

func NewWeatherService(weatherRepo weather_repository.WeatherRepo) *weatherService {
	return &weatherService{
		weatherRepo: weatherRepo,
	}
}

func (w *weatherService) CreateWeather(weatherPayload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage) {
	err := helpers.ValidateStruct(weatherPayload)

	if err != nil {
		return nil, err
	}

	weather := entity.Weather{
		Water: weatherPayload.Water,
		Wind: weatherPayload.Wind,
	}

	weather.DetermineStatus()

	weatherData, err := w.weatherRepo.CreateWeather(weather)

	if err != nil {
		return nil, err
	}

	response := &dto.WeatherResponse{
		Water: weatherData.Water,
		Wind: weatherData.Wind,
		WaterStatus: weatherData.WaterStatus,
		WindStatus: weatherData.WindStatus,
		StatusCode: http.StatusOK,
	}

	return response, nil
}

func (w *weatherService) GetLastWeather() (*dto.WeatherResponse, errs.ErrMessage) {
	weatherData, err := w.weatherRepo.GetLastWeather()

	if err != nil {
		return nil, err
	}

	response := &dto.WeatherResponse{
		Water: weatherData.Water,
		Wind: weatherData.Wind,
		WaterStatus: weatherData.WaterStatus,
		WindStatus: weatherData.WindStatus,
		StatusCode: http.StatusOK,
	}

	return response, nil
}