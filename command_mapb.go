package main

import (
	"fmt"
)

func commandMapBack(config *Config) error {
	locationArea, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		fmt.Println(err)
	}
	config.nextLocationsURL = &locationArea.Next
	config.prevLocationsURL = &locationArea.Previous

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)

	}

	return nil
}
