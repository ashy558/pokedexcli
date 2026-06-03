package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *pokeapi.Configuration, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("one additional argument is required")
	}
	if err := cfg.Pokedex.Inspect(pokemonName); err != nil {
		return fmt.Errorf("pokedex inspect: %w", err)
	}
	return nil
}
