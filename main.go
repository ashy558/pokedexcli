package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/pokecache"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func main() {
	pokeClient := http.DefaultClient
	nextLocationAreaURL, err := url.JoinPath(pokeapi.BaseURL, "location-area")
	if err != nil {
		fmt.Printf("pokedexcli: url error: %s", err)
	}
	newCache := pokecache.NewCache(60 * time.Second)

	cfg := &poketypes.Configuration{
		Client:   pokeClient,
		Next:     nextLocationAreaURL,
		Previous: "",
		Cache:    &newCache,
	}
	startRepl(cfg)
}
