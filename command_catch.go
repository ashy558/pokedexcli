package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *pokeapi.Configuration, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("one additional argument is required")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if err := pokeapi.AttemptCapture(cfg, pokemon); err != nil {
		return fmt.Errorf("attempt capture: %w", err)
	}
	return nil
}
