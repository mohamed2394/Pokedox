package main

import (
	"fmt"
)

func printPrompt() {
	fmt.Print("pokedox", "> ")
}
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

func handleCmd(text string) {
	handleInvalidCmd(text)
}
