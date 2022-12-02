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
	curl  "https://api.openweathermap.org/data/2.5/weather?q= " + city + " + &appid=" + APIkey
	//curl "https://localhost:3000?id=abcdefghi"
	urlParse, err := url.Parse(curl)
	resp, err := http.Get(urlParse.String())

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	fmt.Println((s))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := r.URL.Query().Get("temp")
		fmt.Println("temp =>", temp)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pressure := r.URL.Query().Get("pressure")
		fmt.Println("Pressure =>", pressure)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		humidity := r.URL.Query().Get("humidity")
		fmt.Println("Humidity=>", humidity)
	})
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//temp := r.URL.Query().Get("temp")*/
	//fmt.Println(s.temp)
	//pressure := r.URL.Query().Get("pressure")
	//fmt.Println(s.pressure)
	//humidity := r.URL.Query().Get("humidity")
	//fmt.Println(s.humidity)

}
