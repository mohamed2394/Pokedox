package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"clear": {
		name:        "clear",
		description: "Clears the Terminal",
		callback:    clearScreen,
	},
	"map": {
		name:        "map",
		description: "displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "displays the previous 20 locations in the Pokemon world",
		callback:    commandMapB,
	},
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()

	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			command.callback()
		} else if strings.EqualFold("exit", text) {
			return
		} else {
			handleCmd(text)
		}
		printPrompt()
	}
	fmt.Println()
}
