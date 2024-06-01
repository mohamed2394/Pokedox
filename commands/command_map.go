package commands

import (
	"fmt"
	"log"
	"pokedox/config"
)

func CommandMap(cfg *config.Config, parameter ...string) error {
	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaUrl = resp.Next
	cfg.PreviousLocationAreaUrl = resp.Previous

	return nil
}

func CommandMapB(cfg *config.Config, parameter ...string) error {
	if cfg.PreviousLocationAreaUrl == nil {
		return fmt.Errorf("you are in the first instance, no previous location areas")
	}

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PreviousLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaUrl = resp.Next
	cfg.PreviousLocationAreaUrl = resp.Previous
	return nil
}

func CommandExplore(cfg *config.Config, parameters ...string) error {
	// Check if exactly one parameter is provided
	if len(parameters) != 1 {
		return fmt.Errorf("explore command requires exactly one parameter")
	}

	// Extract the single parameter
	area := parameters[0]

	// Call the ListAreaPokemones function with the parameter
	resp, err := cfg.PokeapiClient.ListAreaPokemones(area)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exploring pastoria-city-area....")
	if len(resp.PokemonEncounters) != 0 {
		fmt.Println("Found Pokemones :")
		for _, encounter := range resp.PokemonEncounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemones Found here .")

	}
	return nil
}
