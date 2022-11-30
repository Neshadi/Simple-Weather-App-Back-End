package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// weather struct
type WeatherResponse struct {
	Name string `json: "Name"`
	Main Main   `json: "Main"`
}

type Main struct {
	temp     float64 `json: "temp"`
	pressure float64 `json: "pressure"`
	humidity float64 `json: "Main"`
}

func main() {

	fmt.Println("Enter the country name")
	var city string
	fmt.Scanln(&city)
	APIkey := "75c150e42d3efb8adc8d819361985205"
	URL := "https://api.openweathermap.org/data/2.5/weather?q= " + city + " + &appid=" + APIkey
	urlParse, err := url.Parse(URL)
	resp, err := http.Get(urlParse.String())

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}