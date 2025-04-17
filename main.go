package main

import (
	"github.com/hvilander/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	caughtPokemon := make(map[string]pokeapi.Pokemon)
	config := &Config{
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemon,
	}

	startRepl(config)
}
