package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Data struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"wheather"`
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

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	Feelslike float64 `json:"feelslike"`
	Tempmin   float64 `json:"tempmin"`
	Temp_max  float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"haze"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Sys struct {
	Type    int64  `json:"type"`
	Id      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

func main() {

	var d *Data
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://community-open-weather-map.p.rapidapi.com/weather", nil)

	req.Header.Add("X-RapidAPI-Host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "d9a4e2d051mshb30a118837bacc6p13b99bjsn741777ede9ea")
	q := url.Values{}

	q.Add("q", "Hyderabad,India")
	q.Add("lat", "17.5382624074844")
	q.Add("lon", "78.5070977778691")
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
		log.Println(err)
	}

	fmt.Println(d)
}
