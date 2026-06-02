package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *configuration) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(cfg.Previous)
	if err != nil {
		return fmt.Errorf("mapb: http get error: %w", err)
	}
	defer res.Body.Close()
	var out pokeapi.LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&out); err != nil {
		return fmt.Errorf("mapb: body decoder error: %w", err)
	}
	cfg.Previous = out.Previous
	cfg.Next = out.Next
	for _, locationArea := range out.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
