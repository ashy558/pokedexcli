package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandMapb(cfg *poketypes.Configuration, args string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	err := pokeapi.PrintLocationAreas(cfg, cfg.Previous)
	if err != nil {
		return fmt.Errorf("PrintLocationAreas: %w", err)
	}
	return nil
}
