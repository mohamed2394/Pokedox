package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandMap() error {

	return nil
}

func commandMapB() error {

	return nil
}

func main() {
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
