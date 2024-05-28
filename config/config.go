package config

import "pokedox/internal/pokeapi"

type Config struct {
	PokeapiClient           pokeapi.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}
