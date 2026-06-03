package pokeapi

import (
	"fmt"
	"sync"
)

type Pokedex struct {
	History map[string]PokemonSpecies
	Lock    sync.RWMutex
}

func (p *Pokedex) Add(cfg *Configuration, pokemonName string) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	species, err := getPokemonSpecies(cfg, pokemonName)
	if err != nil {
		return fmt.Errorf("get pokemon species: %w", err)
	}
	p.History[pokemonName] = species
	return nil
}

func (p *Pokedex) List() error {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	if len(p.History) == 0 {
		fmt.Println("The Pokedex is empty!")
		return nil
	}
	for name, _ := range p.History {
		fmt.Printf("- %s", name)
	}
	return nil
}
