package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandMapb(cfg *poketypes.Configuration) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	err := pokeapi.PrintLocationAreaResponse(cfg.Previous, cfg)
	if err != nil {
		return fmt.Errorf("mapb: PrintLocationAreaResponse: %w", err)
	}
	return nil
}
