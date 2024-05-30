package main

import (
	"pokedox/config"
	"pokedox/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := &config.Config{
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
