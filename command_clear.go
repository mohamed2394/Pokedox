package main

import "github.com/inancgumus/screen"

func clearScreen(cfg *config) error {
	screen.Clear()
	screen.MoveTopLeft()
	return nil
}
