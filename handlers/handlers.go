package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type Response struct {
	Time         string  `json:"time"`
	Temp         float64 `json:"temp"`
	ChanceOfRain float64 `json:"chance_of_rain"`
	Condition    string  `json:"condition"`
	Icon         string  `json:"icon"`
}

type ResponseNow struct {
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Temp      float64 `json:"temperature"`
	Condition string  `json:"condition"`
	Icon      string  `json:"icon"`
}

func GetCityFuture(context *gin.Context) {
	city, isHere := context.Params.Get("city")
	if !isHere {
		context.Status(http.StatusBadRequest)
	}
	req, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=b3eac86bb01f44ce9ae101830232208&q=" + city + "&days=1&aqi=no&alerts=no")
	if err != nil {
		context.JSON(http.StatusBadRequest, "Wrong city")
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, "error while reading req.Body")
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		context.JSON(http.StatusBadRequest, fmt.Sprintf("Error while unmarshalling: %s", err))
	}
	hours := weather.Forecast.Forecastday[0].Hour
	allResponses := []Response{}
	for _, hour := range hours {
		if !time.Unix(hour.TimeEpoch, 0).Before(time.Now()) && time.Unix(hour.TimeEpoch, 0).Sub((time.Now().Add(time.Hour*6))) <= time.Hour*11 {
			var response Response
			response.Time = time.Unix(hour.TimeEpoch, 0).Format("15:04")
			response.Temp = hour.TempC
			response.ChanceOfRain = hour.ChanceOfRain
			response.Condition = hour.Condition.Text
			response.Icon = hour.Condition.Icon
			allResponses = append(allResponses, response)
		}
	}
	location := weather.Location
	current := weather.Current
	nowResponse := ResponseNow{
		City:      location.Name,
		Country:   location.Country,
		Temp:      current.TempC,
		Condition: current.Condition.Text,
		Icon:      current.Condition.Icon,
	}
	context.HTML(http.StatusOK, "index.html", gin.H{
		"AllResponses": allResponses,
		"ResponseNow":  nowResponse,
	})
}
