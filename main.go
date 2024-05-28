package main

import "pokedox/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
