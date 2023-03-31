package weather_test

import (
	"testing"

	"github.com/DuncanStewart1234/Project-Group-55/Golang-Angular/src/server/weather"
)

func TestGetWeather(t *testing.T) {
	got := weather.GetWeather("Gainesville", "f", "en")

	if got == `Current weather for :
	Conditions:
	Now:         0 imperial
	High:        0 imperial
	Low:         0 imperial` {
		t.Errorf("Current weather not returned, fatal error occured.")
	}
}

func TestGetForecast(t *testing.T) {
	got := weather.GetForecast("Gainesville", "f", "en")

	if got == "Weather Forecast for :" {
		t.Errorf("Forecast not returned, fatal error occured.")
	}
}

func TestGetWeatherMultiWordCity(t *testing.T) {
	got := weather.GetWeather("Los Angeles", "f", "en")

	if got == `Current weather for :
	Conditions:
	Now:         0 imperial
	High:        0 imperial
	Low:         0 imperial` {
		t.Errorf("Weather for city not returned, fatal error occured.")
	}
}

func TestGetForecastCityWithSpecialCharacters(t *testing.T) {
	got := weather.GetForecast("Cancun", "f", "en")

	if got == "Weather Forecast for :" {
		t.Errorf("Forecast for city not returned, fatal error occured.")
	}
}

func TestGetWeatherBadInput(t *testing.T) {
	got := weather.GetWeather("Ggssf", "f", "en")

	if got == "" {
		t.Errorf("False input not handled, fatal error occured.")
	}
}

func TestGetForecastBadInput(t *testing.T) {
	got := weather.GetForecast("xsafsdf", "f", "en")

	if got == "" {
		t.Errorf("False input not handled, fatal error occured.")
	}
}
