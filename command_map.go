package main

import (
	"fmt"
	"log"
	"pokedox/internal/pokeapi"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Locartion Areas :")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
