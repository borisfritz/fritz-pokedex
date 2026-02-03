package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ReplConfig struct {
	Next *string
	Prev *string
}

func startRepl() {
	cfg := &ReplConfig{
		Next: nil,
		Prev: nil,
	}
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := commands[commandName]
		if ok {
			err := command.callback(cfg)
			if err != nil {
				log.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command.  For commands, use command 'help'.")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
