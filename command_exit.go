package main

import (
	"fmt"
	"os"

	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func commandExit(cfg *poketypes.Configuration) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
