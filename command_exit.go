package main

import (
	"fmt"
	"os"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *pokeapi.Configuration, args string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
