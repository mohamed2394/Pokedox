package main

import (
	"pokedox/config"
	"pokedox/internal/pokeapi"
)

func main() {
	cfg := config.Config{
		PokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
