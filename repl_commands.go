package main

import (
	"fmt"
	"os"
	"sort"
)

//HACK: Constant Variable for BaseURL in case it changes (example v3)
const BaseURL = "https://pokeapi.co/api/v2"

type cliCommand struct {
	name 		string
	description string
	callback    func(*replConfig, ...string) error
}

//NOTE: map of repl commands
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
		"explore": {
			name: "explore <location area>",
			description: "Display the pokemon that live in the specified area.",
			callback: commandExplore,
		},
		"pokedex": {
			name: "pokedex",
			description: "List pokemon in your pokedex.",
			callback: commandPokedex,
		},
		"catch": {
			name: "catch <pokemon>",
			description: "Attempt to catch a <pokemon>.",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon>",
			description: "Inspect stats of a <pokemon> in your pokedex",
			callback: commandInspect,
		},
		//NOTE: Data and Debug commands:
		"data": {
			name: "data <command>",
			description: "get data for <command>. commands:\n    -" + colorGreen + "dist" + colorReset + ": Displays pokemon base experiance distribution.",
			callback: commandData,
		},
	}
}

//NOTE: repl commands
func commandExit(cfg *replConfig, args ...string) error {
	fmt.Printf("%vClosing the Pokedex... Goodbye!%v", colorGreen, colorReset)
	os.Exit(0)
	return nil
}

func commandHelp(cfg *replConfig, args ...string) error {
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

func commandMap(cfg *replConfig, args ...string) error {
	var url string
	if cfg.Next == nil {
		url = BaseURL + "/location-area"
	} else {
		url = *cfg.Next
	}
	locationData, err := cfg.Client.GetLocationAreaBatch(url)
	if err != nil {
		return fmt.Errorf("GetLocationAreaData(%v) failed: %w", url, err)
	}
	locationData.PrintNames()
	cfg.Next = locationData.Next
	cfg.Prev = locationData.Previous
	return nil
}

func commandMapb(cfg *replConfig, args ...string) error {
	if cfg.Prev == nil {
		fmt.Printf("%vYou have yet to use 'map' or are currently vewing the first page.%v\n", colorRed, colorReset)
		return nil
	}
	locationData, err := cfg.Client.GetLocationAreaBatch(*cfg.Prev)
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

func commandExplore(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("%vYou must specify a location area to search! Type help for usage.%v", colorRed, colorReset)
		return nil
	}
	if len(args) > 1 {
		fmt.Printf("%vOnly specify one location area! Type help for usage.%v", colorRed, colorReset)
	}
	locationName := args[0]
	url := BaseURL + "/location-area/" + locationName
	locationData, err := cfg.Client.GetLocationAreaEndpoint(url)
	if err != nil {
		return fmt.Errorf("GetLocationAreaEndpoint(%v) failed: %w", url, err)
	}
	locationData.PrintPokemon()	
	return nil
}


func commandPokedex(cfg *replConfig, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.Pokedex {
		fmt.Printf("  - %v\n", v.Name)
	}
	return nil
}

func commandCatch(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("%vYou must specify a pokemon! Type help for usage.%v", colorRed, colorReset)
		return nil
	}
	if len(args) > 1 {
		fmt.Printf("%vOnly specify one pokemon! Type help for usage.%v", colorRed, colorReset)
	}
	pokemon := args[0]
	url := fmt.Sprintf("%v/pokemon/%v", BaseURL, pokemon)
	pokemonData, err := cfg.Client.GetPokemonData(url)
	if err != nil {
		return fmt.Errorf("GetPokemonData(%v) failed: %w", url, err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonData.Name)
	success := pokemonData.AttemptCapture()
	var msg string
	if success {
		msg = fmt.Sprintf("%v was caught!", pokemonData.Name) 
		cfg.Pokedex[pokemonData.Name] = pokemonData
	} else {
		msg = fmt.Sprintf("%v escaped!", pokemonData.Name)
	}
	fmt.Println(msg)
	return nil
}

func commandInspect(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("%vYou must specify a pokemon! Type help for usage.%v", colorRed, colorReset)
		return nil
	}
	if len(args) > 1 {
		fmt.Printf("%vOnly specify one pokemon! Type help for usage.%v", colorRed, colorReset)
	}
	pokemon := args[0]
	pokemonData, ok := cfg.Pokedex[pokemon]
	if ok {
		fmt.Printf("Name: %v\n", pokemonData.Name)
		fmt.Printf("Height: %v\n", pokemonData.Height)
		fmt.Printf("Weight: %v\n", pokemonData.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemonData.Stats {
			fmt.Printf("  -%v: %v\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types: ")
		for _, v := range pokemonData.Types {
			fmt.Printf("  - %v\n", v.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon!")
	}

	return nil
}

//NOTE: data collection and debug commands

func commandData(cfg *replConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("%vNo Data command found. Usage: data <command>.%v", colorRed, colorReset)
		return nil
	}
	if len(args) > 1 {
		fmt.Printf("%vToo many data commands given.  Usage data <command>.%v", colorRed, colorReset)
	}
	commandName := args[0]
	if commandName == "dist" {
		allPokemon := make(map[int]int)
		fmt.Println("Gathering Pokemon Data...")
		for i := 1; i <= 1025; i++ {
			url := fmt.Sprintf("%v/pokemon/%v", BaseURL, i)
			pokemonData, err := cfg.Client.GetPokemonData(url)
			if err != nil {
				return fmt.Errorf("Error getting pokemon data: %w", err)
			}
			allPokemon[pokemonData.ID] = pokemonData.BaseExperience
		}
		dist := make(map[int]int)
		for i := 0; i < 400; i += 10 {
			dist[i] = 0
		}
		for _, v := range allPokemon {
			binStart := (v / 10) * 10
			if binStart >= 0 && binStart < 1000 {
				dist[binStart]++
			}
		}
		keys := make([]int, 0, len(dist))
		for k := range dist {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		fmt.Println("XP Range Distribution:")
		for _, k := range keys {
			fmt.Printf("%4d - %4d: %d pokemon\n", k, k+9, dist[k])
		}
		return nil
	}
	return nil
}


















