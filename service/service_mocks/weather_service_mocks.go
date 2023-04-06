package service_mocks

import (
	"github.com/fydhfzh/assignment-3/dto"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/service"
)

type weatherMockService struct {}

var (
	CreateWeather func(payload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage)
	GetLastWeather func() (*dto.WeatherResponse, errs.ErrMessage)
)

func NewWeatherMockService() service.WeatherService {
	return &weatherMockService{}
}

func (w *weatherMockService) CreateWeather(payload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage) {
	return CreateWeather(payload)
}

func (w *weatherMockService) GetLastWeather() (*dto.WeatherResponse, errs.ErrMessage) {
	return GetLastWeather()
}

