package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Printf(
		"Welcome! These are the available commands:")
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Exits the Command line ")

	return nil
}
