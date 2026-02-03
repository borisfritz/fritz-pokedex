package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl(cfg *replConfig) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	//NOTE: print startup text
	commands["help"].callback(cfg)

	//NOTE: Read-Evaluate-Print Loop (REPL)
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

//NOTE: Helper functions
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
