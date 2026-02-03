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

	commands["help"].callback(cfg)
	for {
		fmt.Printf("%v---Pokedex >%v ", colorYellow, colorReset)
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
				log.Println(colorRed, err, colorReset)
			}
			continue
		} else {
			fmt.Printf("%vUnknown Command.  For commands, use command 'help'.%v\n", colorRed, colorReset)
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
