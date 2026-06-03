package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *pokeapi.Configuration, args string) error {
	err := pokeapi.PrintLocationAreas(cfg, cfg.Next)
	if err != nil {
		return fmt.Errorf("PrintLocationAreas: %w", err)
	}
	return nil
}
