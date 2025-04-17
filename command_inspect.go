package main

import (
	"errors"
	"fmt"
)

func command_inspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you can only inspect exactly one pokemon")
	}
	pokename := args[0]

	// check if we have caught one
	pokemon, caught := config.caughtPokemon[pokename]
	if caught {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: ", stat.Stat.Name)
			fmt.Printf(" %d\n", stat.BaseStat)
		}
		fmt.Println("Types:")

		for _, ptype := range pokemon.Types {
			fmt.Printf("  - %s\n", ptype.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
