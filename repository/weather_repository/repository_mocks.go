package weather_repository

import (
	"github.com/fydhfzh/assignment-3/entity"
	"github.com/fydhfzh/assignment-3/pkg/errs"
)

type weatherMockRepo struct{}

var (
	CreateWeather func(weather entity.Weather) (*entity.Weather, errs.ErrMessage)
	GetLastWeather func() (*entity.Weather, errs.ErrMessage)
)

func NewWeatherMockRepository() WeatherRepo {
	return &weatherMockRepo{}
}

func (w *weatherMockRepo) CreateWeather(weather entity.Weather) (*entity.Weather, errs.ErrMessage) {
	return CreateWeather(weather)
}

func (w *weatherMockRepo) GetLastWeather() (*entity.Weather, errs.ErrMessage) {
	return GetLastWeather()
}