package weather_pg

import (
	"database/sql"

	"github.com/fydhfzh/assignment-3/entity"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/repository/weather_repository"
)

type weatherPG struct {
	db *sql.DB
}

func NewWeatherPG(db *sql.DB) weather_repository.WeatherRepo {
	return &weatherPG{
		db: db,
	}
}

func (w *weatherPG) CreateWeather(weatherData entity.Weather) (*entity.Weather, errs.ErrMessage) {
	insertWeatherQuery := `
		INSERT INTO weather (water, wind, water_status, wind_status)
		VALUES($1, $2, $3, $4)
		RETURNING water, wind, water_status, wind_status;
	`

	row := w.db.QueryRow(insertWeatherQuery, weatherData.Water, weatherData.Wind, weatherData.WaterStatus, weatherData.WindStatus)

	var weather entity.Weather

	err := row.Scan(&weather.Water, &weather.Wind, &weather.WaterStatus, &weather.WindStatus)

	if err != nil {
		return nil, errs.NewNotFoundError("row not found")
	}

	return &weather, nil
}

func (w *weatherPG) GetLastWeather() (*entity.Weather, errs.ErrMessage) {
	getWeatherQuery := `
		SELECT water, wind, water_status, wind_status
		FROM weather
		ORDER BY created_at DESC
		LIMIT 1;
	`

	row := w.db.QueryRow(getWeatherQuery)

	var weather entity.Weather

	err := row.Scan(&weather.Water, &weather.Wind, &weather.WaterStatus, &weather.WindStatus)

	if err != nil {
		return nil, errs.NewNotFoundError("row not found")
	}

	return &weather, nil
}