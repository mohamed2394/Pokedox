package commands

import (
	"fmt"
	"os"

	"pokedox/config"

	"github.com/inancgumus/screen"
)

func ClearScreen(cfg *config.Config) error {
	screen.Clear()
	screen.MoveTopLeft()
	return nil
}
func CommandExit(cfg *config.Config) error {
	os.Exit(1)
	return nil
}

func CommandHelp(cfg *config.Config) error {
	fmt.Printf(
		"Welcome! These are the available commands:")
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Exits the Command line ")

	return nil
}
