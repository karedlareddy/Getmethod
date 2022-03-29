package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"io/ioutil"
)

type Data struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}
type Sys struct {
	Type    int64  `json:"type"`
	Id      int64  `json:"id"`
	Sunset  int64  `json:"sunset"`
	Sunrise int64  `json:"sunrise"`
	Country string `json:"country"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Weather struct {
	Id          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	Feelslike float64 `json:"feelslike"`
	Tempmin   float64 `json:"tempmin"`
	Temp_max  float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
type Clouds struct {
	All int64 `json:"all"`
}

func main() {

	var d *Data
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://community-open-weather-map.p.rapidapi.com/weather", nil)

	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "ab4722269bmsh56bf6d667ec4f0ap18748fjsnd5ce885c0f22")
	q := url.Values{}

	q.Add("q", "hyderabad,india")
	q.Add("lat", "17.462741745617286")
	q.Add("lon", "78.37850846489877")
	q.Add("lang", "English-en")

	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
