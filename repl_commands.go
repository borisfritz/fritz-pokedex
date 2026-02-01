package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name 			string
	description 	string
	callback func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	if len(commands) == 0 {
		return fmt.Errorf("Commands not found")
	}
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
