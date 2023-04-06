package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fydhfzh/assignment-3/dto"
	"github.com/fydhfzh/assignment-3/pkg/errs"
	"github.com/fydhfzh/assignment-3/service/service_mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWeatherHandler_CreateNewWeather_Success(t *testing.T) {
	payload := dto.WeatherRequest{
		Wind: 18,
		Water: 18,
	}
	
	weatherService := service_mocks.NewWeatherMockService()

	weatherHandler := NewWeatherHandler(weatherService)

	service_mocks.CreateWeather = func(payload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage) {
		response := dto.WeatherResponse{
			Water: 18,
			Wind: 18,
			WaterStatus: "Bahaya",
			WindStatus: "Bahaya",
			StatusCode: http.StatusOK,
		}

		return &response, nil
	}

	jsonByte, err := json.Marshal(payload)

	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "/weather", bytes.NewBuffer(jsonByte))

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/weather", weatherHandler.CreateWeather)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var weatherResponse dto.WeatherResponse

	err = json.Unmarshal(responseBody, &weatherResponse)

	require.Nil(t, err)

	assert.Equal(t, 18, weatherResponse.Water)
	assert.Equal(t, 18, weatherResponse.Wind)
	assert.Equal(t, "Bahaya", weatherResponse.WaterStatus)
	assert.Equal(t, "Bahaya", weatherResponse.WindStatus)
	assert.Equal(t, http.StatusOK, weatherResponse.StatusCode)
}

func TestWeatherHandler_CreateNewWeather_UnprocessableEntityError(t *testing.T) {
	payload := dto.WeatherRequest{
		Wind: 18,
	}
	
	weatherService := service_mocks.NewWeatherMockService()

	weatherHandler := NewWeatherHandler(weatherService)

	service_mocks.CreateWeather = func(payload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage) {
		return nil, errs.NewUnprocessableEntityError("invalid request body")
	}

	jsonByte, err := json.Marshal(payload)

	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "/weather", bytes.NewBuffer(jsonByte))

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/weather", weatherHandler.CreateWeather)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var errorResponse errs.ErrorMessage

	err = json.Unmarshal(responseBody, &errorResponse)

	require.Nil(t, err)

	assert.NotNil(t, errorResponse)
	assert.Equal(t, "invalid request body", errorResponse.Message())
	assert.Equal(t, "INVALID_REQUEST_BODY", errorResponse.Error())
	assert.Equal(t, http.StatusUnprocessableEntity, errorResponse.Status())
}

func TestWeatherHandler_CreateNewWeather_DbError(t *testing.T) {
	payload := dto.WeatherRequest{
		Wind: 18,
		Water: 18,
	}
	
	weatherService := service_mocks.NewWeatherMockService()

	weatherHandler := NewWeatherHandler(weatherService)

	service_mocks.CreateWeather = func(payload dto.WeatherRequest) (*dto.WeatherResponse, errs.ErrMessage) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	jsonByte, err := json.Marshal(payload)

	require.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "/weather", bytes.NewBuffer(jsonByte))

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/weather", weatherHandler.CreateWeather)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var errorResponse errs.ErrorMessage

	err = json.Unmarshal(responseBody, &errorResponse)

	require.Nil(t, err)

	assert.NotNil(t, errorResponse)
	assert.Equal(t, "something went wrong", errorResponse.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", errorResponse.Error())
	assert.Equal(t, http.StatusInternalServerError, errorResponse.Status())
}

func TestWeatherHandler_GetLastWeather_Success(t *testing.T) {
	weatherService := service_mocks.NewWeatherMockService()

	weatherHandler := NewWeatherHandler(weatherService)

	service_mocks.GetLastWeather = func() (*dto.WeatherResponse, errs.ErrMessage) {
		response := dto.WeatherResponse{
			Water: 18,
			Wind: 18,
			WaterStatus: "Bahaya",
			WindStatus: "Bahaya",
			StatusCode: http.StatusOK,
		}

		return &response, nil
	}

	req, err := http.NewRequest(http.MethodGet, "/api/weather", nil)

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.GET("/api/weather", weatherHandler.GetLastWeather)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var weatherResponse dto.WeatherResponse

	err = json.Unmarshal(responseBody, &weatherResponse)

	require.Nil(t, err)

	assert.Equal(t, 18, weatherResponse.Water)
	assert.Equal(t, 18, weatherResponse.Wind)
	assert.Equal(t, "Bahaya", weatherResponse.WaterStatus)
	assert.Equal(t, "Bahaya", weatherResponse.WindStatus)
	assert.Equal(t, http.StatusOK, weatherResponse.StatusCode)
}

func TestWeatherHandler_GetLastWeather_DbError(t *testing.T) {
	weatherService := service_mocks.NewWeatherMockService()

	weatherHandler := NewWeatherHandler(weatherService)

	service_mocks.GetLastWeather = func() (*dto.WeatherResponse, errs.ErrMessage) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	req, err := http.NewRequest(http.MethodGet, "/api/weather", nil)

	require.Nil(t, err)

	rr := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.GET("/api/weather", weatherHandler.GetLastWeather)

	route.ServeHTTP(rr, req)

	result := rr.Result()

	responseBody, err := ioutil.ReadAll(result.Body)

	require.Nil(t, err)

	defer result.Body.Close()

	var errorResponse errs.ErrorMessage

	err = json.Unmarshal(responseBody, &errorResponse)

	require.Nil(t, err)

	assert.NotNil(t, errorResponse)
	assert.Equal(t, "something went wrong", errorResponse.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", errorResponse.Error())
	assert.Equal(t, http.StatusInternalServerError, errorResponse.Status())
}
