package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListAreaPokemones(areaName string) (ExplorePokemonesResponse, error) {
	// Ensure the URL is correctly constructed
	fullUrl := fmt.Sprintf("%s/location-area/%s", baseUrl, areaName)

	if val, ok := c.cache.Get(fullUrl); ok {
		exploreResponse := ExplorePokemonesResponse{}
		err := json.Unmarshal(val, &exploreResponse)
		if err != nil {
			return ExplorePokemonesResponse{}, err
		}

		return exploreResponse, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return ExplorePokemonesResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExplorePokemonesResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return ExplorePokemonesResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExplorePokemonesResponse{}, err
	}

	exploreResp := ExplorePokemonesResponse{}
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return ExplorePokemonesResponse{}, err
	}

	return exploreResp, nil
}
