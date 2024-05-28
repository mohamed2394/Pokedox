package commands

import (
	"fmt"
	"log"
)

func commandMap(cfg config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locartion Areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousLocationAreaUrl == nil {
		return fmt.Errorf("you are in the first  ")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locartion Areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous

	return nil
}
