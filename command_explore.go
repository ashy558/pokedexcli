package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandExplore(cfg *poketypes.Configuration, area string) error {
	if area == "" {
		return fmt.Errorf("one additional argument is required")
	}
	fmt.Printf("Exploring %s...\n", area)
	err := pokeapi.PrintAreaPokemon(cfg, area)
	if err != nil {
		return fmt.Errorf("PrintAreaPokemon: %w", err)
	}
	return nil
}
