package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedox/commands"
	"pokedox/config"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config.Config) error
}

var commands_e = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commands.CommandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.CommandExit,
	},
	"clear": {
		name:        "clear",
		description: "Clears the Terminal",
		callback:    commands.ClearScreen,
	},
	"map": {
		name:        "map",
		description: "displays the names of 20 location areas in the Pokemon world",
		callback:    commands.CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "displays the previous 20 locations in the Pokemon world",
		callback:    commands.CommandMapB,
	},
}

func startRepl(cfg *config.Config) {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()

	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands_e[text]; exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else if strings.EqualFold("exit", text) {
			return
		} else {
			handleCmd(text)
		}
		printPrompt()
	}
	fmt.Println()
}
