package main

import (
	"fmt"
	"os"

	"github.com/ashy558/pokedexcli/internal/pokeapi"
)

func main() {
	cfg, err := pokeapi.NewConfiguration()
	if err != nil {
		fmt.Printf("new configuration: %s", err)
		os.Exit(1)
	}
	startRepl(cfg)
}
