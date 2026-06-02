package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandMap(cfg *poketypes.Configuration, args string) error {
	err := pokeapi.PrintLocationAreas(cfg, cfg.Next)
	if err != nil {
		return fmt.Errorf("PrintLocationAreas: %w", err)
	}
	return nil
}
