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
	callback    func(cfg *config.Config, parameter ...string) error
}

func printPrompt() {
	fmt.Print("pokedox", "> ")
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
	"explore": {
		name:        "explore",
		description: "explore the are for pokemones",
		callback:    commands.CommandExplore,
	},
	"catch": {
		name:        "catch",
		description: "catch a pokemone",
		callback:    commands.CommandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "inspect a pokemone in hand",
		callback:    commands.CommandInspect,
	},
	"pokedox": {
		name:        "pokedox",
		description: "display the caught pokemons in the pokedox",
		callback:    commands.CommandPokedox,
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
		text := reader.Text()
		prompt := strings.Fields(text)
		if len(prompt) == 0 {
			printPrompt()
			continue
		}

		commandName := prompt[0]
		parameters := prompt[1:]

		if command, exists := commands_e[commandName]; exists {
			err := command.callback(cfg, parameters...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if strings.EqualFold("exit", commandName) {
			return
		} else {
			fmt.Println("Unknown command:", commandName)
		}
		printPrompt()
	}
	fmt.Println()
}
