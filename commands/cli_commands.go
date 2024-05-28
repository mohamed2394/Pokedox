package commands

import (
	"fmt"
	"os"

	"github.com/inancgumus/screen"
)

func clearScreen(cfg config) error {
	screen.Clear()
	screen.MoveTopLeft()
	return nil
}
func commandExit(cfg Config) error {
	os.Exit(1)
	return nil
}

func commandHelp(cfg Config) error {
	fmt.Printf(
		"Welcome! These are the available commands:")
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Exits the Command line ")

	return nil
}
