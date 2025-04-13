package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hvilander/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading stdin:", err)
		}

		if len(input) == 0 {
			continue
		}

		commandWord := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		command, exists := getCommands()[commandWord]

		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandWord)
			continue
		}

	}
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(strings.ToLower(text))
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the PREVIOUS 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Given a location name (e.g. `explore pastoria-city-area`) shows a list of found pokemon",
			callback:    commandExplore,
		},
	}
}
