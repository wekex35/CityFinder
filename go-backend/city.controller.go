package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type StringResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ObjectResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Feature `json:"data"`
}

type Feature struct {
	Type     string          `json:"type"`
	Features []GeoJsonFormat `json:"features"`
}

func countryList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := ListResponse{200, "Success", getCountryList()}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Print(w, string(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func countryCity(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var country = req.URL.Query().Get("country")
	print(country)
	res := ListResponse{200, "Success", getCityListByCountry(country)}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Print(w, string(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func searchCityStr(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var query = req.URL.Query().Get("query")
	print(query)
	res := ListResponse{200, "Success", searchCity(query)}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Print(w, string(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func adjacentCity(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var city = req.URL.Query().Get("city")
	print(city)
	res := ObjectResponse{200, "Success", Feature{"FeatureCollection", getAdjacentCityList(city)}}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Print(w, string(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
