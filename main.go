package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := http.DefaultClient
	nextLocationAreaURL, err := url.JoinPath(pokeapi.BaseURL, "location-area")
	if err != nil {
		fmt.Printf("pokedexcli: url error: %s", err)
	}

	cfg := &configuration{Client: pokeClient, Next: nextLocationAreaURL, Previous: ""}
	startRepl(cfg)
}
