package main

import "github.com/inancgumus/screen"

func clearScreen() error {
	screen.Clear()
	screen.MoveTopLeft()
	return nil
}
