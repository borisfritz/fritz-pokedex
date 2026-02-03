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
			description: "Display the first (or next) 20 location areas.",
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
	fmt.Printf("%vClosing the Pokedex... Goodbye!%v", colorGreen, colorReset)
	os.Exit(0)
	return nil
}

func commandHelp(cfg *ReplConfig) error {
	fmt.Printf("%vWelcome to the Pokedex!\n%v---Usage---\n", colorYellow, colorReset)
	commands := getCommands()
	if len(commands) == 0 {
		return fmt.Errorf("Command not found!")
	}
	for _, cmd := range commands {
		fmt.Printf("%v%v%v: %v\n", colorGreen, cmd.name, colorReset, cmd.description)
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
	locationData.PrintNames()
	cfg.Next = locationData.Next
	cfg.Prev = locationData.Previous
	return nil
}

func commandMapb(cfg *ReplConfig) error {
	if cfg.Prev == nil {
		fmt.Printf("%vYou are already on the first page.%v\n", colorRed, colorReset)
		return nil
	}
	locationData, err := pokeapi.GetLocationAreaData(*cfg.Prev)
	if err != nil {
		return fmt.Errorf("GetLocationAreaData(%v) failed: %w", *cfg.Prev, err)
	}
	locationData.PrintNames()
	cfg.Next = locationData.Next
	if locationData.Previous != nil {
		cfg.Prev = locationData.Previous
	} else {
		cfg.Prev = nil
	}
	return nil
}
