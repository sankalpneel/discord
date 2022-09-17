package temperature

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sankalpneel/discord/config"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func Query(city string) string {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?appid=" + config.OpenWeatherMapApiKey + "&q=" + city)

	if err != nil {
		return "Something Went Wrong!!!"
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return "Something Went Wrong!!!"
	}

	celc := (d.Main.Kelvin - 273.15)
	s := fmt.Sprintf("%.2f", celc)
	return s
}
