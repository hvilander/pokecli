package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	// get location name from args
	if len(args) != 1 {
		return errors.New("You must provide exactly one location name to explore")
	}

	locationName := args[0]

	locationArea, err := config.pokeapiClient.LocationByName(locationName)
	if err != nil {
		return fmt.Errorf("command explore failed: %w", err)
	}

	fmt.Printf("Exploring %s\n", locationName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf("  - %s\n", encounter.Pokemon.Name)

	}
	return nil
}
