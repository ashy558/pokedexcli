package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandMap(cfg *poketypes.Configuration) error {
	err := pokeapi.PrintLocationAreaResponse(cfg.Next, cfg)
	if err != nil {
		return fmt.Errorf("map: PrintLocationAreaResponse: %w", err)
	}
	return nil
}
