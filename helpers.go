package main

import (
	"fmt"
	"strings"
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
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
