package pokeapi

import (
	"net/http"
	"pokedox/internal/pokecache"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
