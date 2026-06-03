package pokeapi

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/ashy558/pokedexcli/internal/pokecache"
)

type Configuration struct {
	Client   *http.Client
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  Pokedex
}

func NewConfiguration() (*Configuration, error) {
	nextLocationAreaURL, err := url.JoinPath(BaseURL, "location-area")
	if err != nil {
		return nil, fmt.Errorf("pokedexcli: url error: %s", err)
	}
	newCache := pokecache.NewCache(60 * time.Second)
	return &Configuration{
		Client:   http.DefaultClient,
		Next:     nextLocationAreaURL,
		Previous: "",
		Cache:    &newCache,
		Pokedex: Pokedex{
			History: map[string]PokemonSpecies{},
			Lock:    sync.RWMutex{},
		},
	}, nil
}
