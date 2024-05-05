package main

import "os"

func commandExit(cfg *config) error {
	os.Exit(1)
	return nil
}
