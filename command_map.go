package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationArea, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
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
