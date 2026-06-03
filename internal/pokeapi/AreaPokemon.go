package pokeapi

import (
	"fmt"
)

func PrintAreaPokemon(cfg *Configuration, area string) error {
	locationArea, err := getLocationArea(cfg, area)
	if err != nil {
		return fmt.Errorf("getLocationArea: %w", err)
	}
	encounters := locationArea.PokemonEncounters
	if len(encounters) == 0 {
		fmt.Println("No Pokemon could be found.")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
