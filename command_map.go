package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationArea, err := getLocationArea(config)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)

	}

	return nil
}
