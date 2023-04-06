package dto

type WeatherRequest struct {
	Water int `json:"water" valid:"required~water cannot be empty"`
	Wind  int `json:"wind" valid:"required~wind cannot be empty"`
}

type WeatherResponse struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
	StatusCode  int    `json:"status_code"`
}