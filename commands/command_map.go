package commands

import (
	"fmt"
	"log"
	"pokedox/config"
)

func CommandMap(cfg *config.Config) error {

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locartion Areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaUrl = resp.Next
	cfg.PreviousLocationAreaUrl = resp.Previous

	return nil
}

func CommandMapB(cfg *config.Config) error {
	if cfg.PreviousLocationAreaUrl == nil {
		return fmt.Errorf("you are in the first  ")
	}

	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PreviousLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locartion Areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.NextLocationAreaUrl = resp.Next
	cfg.PreviousLocationAreaUrl = resp.Previous

	return nil
}
