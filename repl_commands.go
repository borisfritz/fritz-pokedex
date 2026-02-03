package main

import (
	"fmt"
	"os"
	"github.com/borisfritz/fritz-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name 			string
	description 	string
	callback func(*ReplConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Display first 20 location areas.  Use again to display next 20 items.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the previous 20 location areas.",
			callback: commandMapb,
		},
	}
}

func commandExit(cfg *ReplConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *ReplConfig) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	if len(commands) == 0 {
		return fmt.Errorf("Command not found!")
	}
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *ReplConfig) error {
	var url string
	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.Next
	}
	locationData, err := pokeapi.GetLocationAreaData(url)
	if err != nil {
		return fmt.Errorf("GetLocationAreaData(%v) failed: %w", url, err)
	}
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = locationData.Next
	cfg.Prev = locationData.Previous
	return nil
}

func commandMapb(cfg *ReplConfig) error {
	if cfg.Prev == nil {
		fmt.Println("You are on the first page.")
		return nil
	}
	locationData, err := pokeapi.GetLocationAreaData(*cfg.Prev)
	if err != nil {
		return fmt.Errorf("GetLocationAreaData(%v) failed: %w", *cfg.Prev, err)
	}
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	cfg.Next = locationData.Next
	if locationData.Previous != nil {
		cfg.Prev = locationData.Previous
	} else {
		cfg.Prev = nil
	}
	return nil
}
