package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Converting the JSON into Go struct
type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	fmt.Println("==============================================\n=                                             =\n============The Weather Forecast for Today :)==============\n=                                             =\n==============================================")

	// ASCII Art
	asciiArt := `
		------               _____
	   /      \ ___\     ___/    ___
	--/-  ___  /    \/  /  /    /   \
   /     /           \__     //_     \
  /                     \   / ___     |
  |           ___       \/+--/        /
   \__           \       \           /
	  \__                 |          /
	 \     /____      /  /       |   /
	  _____/         ___       \/  /\
		   \__      /      /    |    |
		 /    \____/   \       /   //
	 // / / // / /\    /-_-/\//-__-
	  /  /  // /   \__// / / /  //
	 //   / /   //   /  // / // /
	  /// // / /   /  //  / //
   //   //       //  /  // / /
	 / / / / /     /  /    /
`
	fmt.Println(asciiArt)

	// by default the q value is stirling
	q := "Default Location"

	// but the user can also search for the city they would like to know the weather forecast for
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	// An endpoint to make a GET request for the weather forecast
	res, err := http.Get("Please add your weather / weather forecast API here :)" + q + "")
	// Handling error
	if err != nil {
		panic(err)
	}

	// Close the body of the response once the main function finishes executing
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("The API is not working!")
	}

	// Reading the body of the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	// Iterating over the hours in the forecast
	for _, hour := range weather.Forecast.Forecastday[0].Hour {
		date := time.Unix(hour.TimeEpoch, 0)

		fmt.Printf("%s - %.0fC, %.0f%% chance of rain, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		// if hour.ChanceOfRain < 40 {
		// 	fmt.Print(message)
		// } else {
		// 	color.Gray16(message)
		// }

	}

}
