// Examples:
//
//	go run weather.go                             -> displays help message
//	go run weather.go -w Philadelphia -u f -l en  -> Philadelphia, fahrenheit, English
//	go run weather.go -w here -u f -l ru          -> Current location, fahrenheit, Russian
//	go run weather.go -w Dublin -u c -l fi        -> Dublin, celsius, Finnish
//	go run weather.go -w "Las Vegas" -u k -l es   -> Las Vegas, kelvin, Spanish
package weather

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	owm "github.com/briandowns/openweathermap" // "owm" for easier use
)

// URL is a constant that contains where to find the user IP info
const URL = "http://ip-api.com/json"

// template used for output
const weatherTemplate = `Current weather for {{.Name}}:
      Conditions: {{range .Weather}} {{.Description}} {{end}}
      Now:         {{.Main.Temp}} {{.Unit}}
      High:        {{.Main.TempMax}} {{.Unit}}
      Low:         {{.Main.TempMin}} {{.Unit}}
  `

const forecastTemplate = `Weather Forecast for {{.City.Name}}:
  {{range .List}}Date & Time: {{.DtTxt}}
  Conditions:  {{range .Weather}}{{.Main}} {{.Description}}{{end}}
  Temp:        {{.Main.Temp}}
  High:        {{.Main.TempMax}}
  Low:         {{.Main.TempMin}}
  {{end}}
  `

// Pointers to hold the contents of the flag args.
var (
	whereFlag = flag.String("w", "", "Location to get weather.  If location has a space, wrap the location in double quotes.")
	unitFlag  = flag.String("u", "", "Unit of measure to display temps in")
	langFlag  = flag.String("l", "", "Language to display temps in")
	whenFlag  = flag.String("t", "current", "current | forecast")
)

// Data will hold the result of the query to get the IP
// address of the caller.
type Data struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}

// getLocation will get the location details for where this
// application has been run from.
func getLocation() (*Data, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	r := &Data{}
	if err = json.Unmarshal(result, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// getCurrent gets the current weather for the provided
// location in the units provided.
func getCurrent(location, units, lang string) (*owm.CurrentWeatherData, error) {
	w, err := owm.NewCurrent(units, lang, os.Getenv("OWM_API_KEY"))
	if err != nil {
		return nil, err
	}
	w.CurrentByName(location)
	return w, nil
}
func getForecast5(location, units, lang string) (*owm.Forecast5WeatherData, error) {
	w, err := owm.NewForecast("5", units, lang, os.Getenv("OWM_API_KEY"))
	if err != nil {
		return nil, err
	}
	w.DailyByName(location, 5)
	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	return forecast, err
}

func main() {
	os.Setenv("OWM_API_KEY", "53a36ef9cff59aed0b8d724a358d9861")
	flag.Parse()

	// If there's any funkiness with cli args, tuck and roll...
	if len(*whereFlag) <= 1 || len(*unitFlag) != 1 || len(*langFlag) != 2 || len(*whenFlag) <= 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Process request for location of "here"
	if strings.ToLower(*whereFlag) == "here" {
		loc, err := getLocation()
		if err != nil {
			log.Fatalln(err)
		}
		w, err := getCurrent(loc.City, *unitFlag, *langFlag)
		if err != nil {
			log.Fatalln(err)
		}
		tmpl, err := template.New("weather").Parse(weatherTemplate)
		if err != nil {
			log.Fatalln(err)
		}

		// Render the template and display
		err = tmpl.Execute(os.Stdout, w)
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}

	if *whenFlag == "current" {
		// Process request for the given location
		w, err := getCurrent(*whereFlag, *unitFlag, *langFlag)
		if err != nil {
			log.Fatalln(err)
		}
		tmpl, err := template.New("weather").Parse(weatherTemplate)
		if err != nil {
			log.Fatalln(err)
		}
		// Render the template and display
		if err := tmpl.Execute(os.Stdout, w); err != nil {
			log.Fatalln(err)
		}
	} else { //forecast
		w, err := getForecast5(*whereFlag, *unitFlag, *langFlag)
		if err != nil {
			log.Fatalln(err)
		}
		tmpl, err := template.New("forecast").Parse(forecastTemplate)
		if err != nil {
			log.Fatalln(err)
		}
		// Render the template and display
		if err := tmpl.Execute(os.Stdout, w); err != nil {
			log.Fatalln(err)
		}
	}
}
