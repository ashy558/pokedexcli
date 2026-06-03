package pokeapi

import (
	"fmt"
	"sync"
)

type Pokedex struct {
	History map[string]Pokemon
	Lock    sync.RWMutex
}

func (p *Pokedex) Add(cfg *Configuration, pokemonName string) error {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	pokemon, err := getPokemon(cfg, pokemonName)
	if err != nil {
		return fmt.Errorf("get pokemon: %w", err)
	}
	p.History[pokemonName] = pokemon
	return nil
}

func (p *Pokedex) List() error {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	if len(p.History) == 0 {
		fmt.Println("The Pokedex is empty!")
		return nil
	}
	for name := range p.History {
		fmt.Printf("- %s", name)
	}
	return nil
}

func (p *Pokedex) Inspect(pokemonName string) error {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	if pokemon, ok := p.History[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.getName())
		fmt.Printf("Height: %d\n", pokemon.getHeight())
		fmt.Printf("Weight: %d\n", pokemon.getWeight())
		fmt.Println("Stats:")
		stats := pokemon.getStats()
		fmt.Printf("  -hp: %d\n", stats["hp"])
		fmt.Printf("  -attack: %d\n", stats["attack"])
		fmt.Printf("  -defense: %d\n", stats["defense"])
		fmt.Printf("  -special-attack: %d\n", stats["special-attack"])
		fmt.Printf("  -special-defense: %d\n", stats["special-defense"])
		fmt.Printf("  -speed: %d\n", stats["speed"])
		fmt.Println("Types:")
		types := pokemon.getTypes()
		for _, pokemonType := range types {
			fmt.Printf("  - %s\n", pokemonType)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil
}
