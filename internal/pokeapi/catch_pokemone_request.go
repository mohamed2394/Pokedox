package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemone(pokemoneName string) (CatchPokemoenResponse, error) {
	// Ensure the URL is correctly constructed
	fullUrl := fmt.Sprintf("%s/pokemon/%s", baseUrl, pokemoneName)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return CatchPokemoenResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return CatchPokemoenResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return CatchPokemoenResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchPokemoenResponse{}, err
	}

	catchResponse := CatchPokemoenResponse{}
	err = json.Unmarshal(data, &catchResponse)
	if err != nil {
		return CatchPokemoenResponse{}, err
	}

	return catchResponse, nil
}
