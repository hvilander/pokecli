package main

import (
	"fmt"
)

func command_pokedex(config *Config, _ ...string) error {

	if len(config.caughtPokemon) == 0 {
		fmt.Println("You have no pokemon in your pokedex")
		fmt.Println("Catch pokemon to fill up your pokedex")
		fmt.Println("Gotta Catch Them ALL")
		return nil
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
