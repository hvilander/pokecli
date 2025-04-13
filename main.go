package main

import (
	"github.com/hvilander/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
