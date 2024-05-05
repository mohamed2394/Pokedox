package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("Bad status code : %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasResponse{}, err

	}

	locationAreas := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResponse{}, err

	}

	return locationAreas, nil
}
