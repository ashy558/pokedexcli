package main

import (
	"fmt"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *pokeapi.Configuration, args string) error {
	if err := cfg.Pokedex.List(); err != nil {
		return fmt.Errorf("pokedex list: %w", err)
	}
	return nil
}
