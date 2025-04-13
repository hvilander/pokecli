package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://mholt.github.io/json-to-go/
type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaByUrl(url string, config *Config) (Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var locationArea Location

	if err = json.Unmarshal(data, &locationArea); err != nil {
		return Location{}, fmt.Errorf("error unmarshalling data: %w", err)
	}

	config.Next = locationArea.Next
	config.Previous = locationArea.Previous

	return locationArea, nil

}

func getLocationArea(config *Config) (Location, error) {
	defaultURL := "https://pokeapi.co/api/v2/location-area/"
	var url string

	if config.Next == "" {
		url = defaultURL
	} else {
		url = config.Next
	}

	return getLocationAreaByUrl(url, config)
}

func getPreviousLocationArea(config *Config) (Location, error) {
	if config.Previous == "" {
		return Location{}, fmt.Errorf("No previous supplied to getPreviousLocationArea")
	}

	return getLocationAreaByUrl(config.Previous, config)
}
