package main

import (
	"fmt"
)

func commandMapBack(config *Config) error {
	locationArea, err := getPreviousLocationArea(config)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)

	}

	return nil
}
