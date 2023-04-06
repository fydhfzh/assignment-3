package weather_repository

import (
	"github.com/fydhfzh/assignment-3/entity"
	"github.com/fydhfzh/assignment-3/pkg/errs"
)

type WeatherRepo interface {
	CreateWeather(weatherData entity.Weather) (*entity.Weather, errs.ErrMessage)
	GetLastWeather() (*entity.Weather, errs.ErrMessage)
}