package commands

import (
	"fmt"
	"log"
	"math/rand"
	"pokedox/config"
	"pokedox/internal/pokeapi"
	"time"
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
	if len(parameters) != 1 {
		return fmt.Errorf("explore command requires exactly one parameter")
	}

	area := parameters[0]

	resp, err := cfg.PokeapiClient.ListAreaPokemones(area)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exploring %s....\n", resp.Name)
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
func getRandomBinary(difficulty int) int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random float between 0 and 1
	randomValue := rand.Float64()

	// Calculate the threshold based on difficulty
	// Higher difficulty means a higher probability of returning 0
	// Assuming difficulty ranges from 1 to 10
	threshold := float64(difficulty) / 100.0

	if randomValue < threshold {
		return 0
	} else {
		return 1
	}
}

func CommandCatch(cfg *config.Config, parameters ...string) error {
	if len(parameters) != 1 {
		return fmt.Errorf("specify the pokemone you want to catch")
	}

	pokemoneName := parameters[0]

	resp, err := cfg.PokeapiClient.CatchPokemone(pokemoneName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Catching %s\n", resp.Name)
	pokedox := map[string]pokeapi.Pokemon{}
	if _, ok := pokedox[resp.Name]; ok {
		fmt.Printf("%s already in Pokedox \n", resp.Name)
		return nil
	}
	res := getRandomBinary(resp.BaseExperience)
	if res == 1 {
		fmt.Printf("%s Caught\n", resp.Name)
		pokedox[resp.Name] = pokeapi.Pokemon{
			Name:  resp.Name,
			Infos: resp,
		}
		return nil
	} else {
		fmt.Printf("Oups, couln't catch %s\n", resp.Name)
		return nil
	}
}
