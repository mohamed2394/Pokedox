package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemone(pokemoneName string) (Pokemon, error) {
	// Ensure the URL is correctly constructed
	fullUrl := fmt.Sprintf("%s/pokemon/%s", baseUrl, pokemoneName)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	catchResponse := Pokemon{}
	err = json.Unmarshal(data, &catchResponse)
	if err != nil {
		return Pokemon{}, err
	}

	return catchResponse, nil
}
