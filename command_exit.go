package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *configuration) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
