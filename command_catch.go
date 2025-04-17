package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func command_catch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you can only try to catch one at a time")
	}

	pokename := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokename)

	// get info about pokemon
	pokemon, err := config.pokeapiClient.PokemonByName(pokename)
	if err != nil {
		return fmt.Errorf("error retrieving pokemon: %w", err)
	}

	// pikachu is a base of 112 lets catch a pika 1/3 times
	// mewtwo is a 220, 255 is the highest I can find
	catch := rand.Intn(300) > pokemon.BaseExperience

	if catch {
		// add to user's pokedex
		config.caughtPokemon[pokemon.Name] = pokemon

		// print caught message
		fmt.Println(pokemon.Name + " was caught!")

	} else {
		fmt.Println(pokemon.Name + " escaped!")

	}

	return nil
}
