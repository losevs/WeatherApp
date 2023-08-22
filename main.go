package main

import (
	"log"
	"losevs/handlers"

	"github.com/gin-gonic/gin"
)

// type Weather struct {
// 	Location struct {
// 		Name    string `json:"name"`
// 		Country string `json:"country"`
// 	} `json:"location"`
// 	Current struct {
// 		TempC     float64 `json:"temp_c"`
// 		Condition struct {
// 			Text string `json:"text"`
// 		} `json:"condition"`
// 	} `json:"current"`
// 	Forecast struct {
// 		Forecastday []struct {
// 			Hour []struct {
// 				TimeEpoch int64   `json:"time_epoch"`
// 				TempC     float64 `json:"temp_c"`
// 				Condition struct {
// 					Text string `json:"text"`
// 					//Icon string `json:"icon"`
// 				} `json:"condition"`
// 				ChanceOfRain float64 `json:"chance_of_rain"`
// 			} `json:"hour"`
// 		} `json:"forecastday"`
// 	} `json:"forecast"`
// }

func main() {
	// 	query := "Moscow"
	// 	if len(os.Args) >= 2 {
	// 		query = os.Args[1]
	// 	}
	// 	req, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=b3eac86bb01f44ce9ae101830232208&q=" + query + "&days=1&aqi=no&alerts=no")
	// 	if err != nil {
	// 		panic(err) //!
	// 	}
	// 	defer req.Body.Close()
	// 	if req.StatusCode != 200 {
	// 		panic("Weather App isn't available") //!
	// 	}
	// 	body, err := io.ReadAll(req.Body)
	// 	if err != nil {
	// 		panic(err) //!
	// 	}
	// 	var weather Weather
	// 	err = json.Unmarshal(body, &weather)
	// 	if err != nil {
	// 		panic(err) //!
	// 	}
	// 	location := weather.Location
	// 	current := weather.Current
	// 	hours := weather.Forecast.Forecastday[0].Hour
	// 	fmt.Printf("%s, %s, %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)
	// 	for _, hour := range hours {
	// 		if !time.Unix(hour.TimeEpoch, 0).Before(time.Now()) {
	// 			fmt.Printf("%s - %.0fC, %.0f%%, %s\n", time.Unix(hour.TimeEpoch, 0).Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)
	// 		}
	// 	}
	router := gin.Default()
	router.GET("/weather/now/:city", handlers.GetCityNow)
	router.GET("/weather/:city", handlers.GetCityFuture)
	log.Fatalln(router.Run(":80"))
}
