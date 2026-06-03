package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func getPokemon(cfg *Configuration, pokemon string) (out Pokemon, err error) {
	pokemonURL, err := url.JoinPath(BaseURL, "/pokemon/", pokemon)
	if err != nil {
		return out, fmt.Errorf("url join path: %w", err)
	}
	data, err := getWithCache(cfg, pokemonURL)
	if err != nil {
		return out, fmt.Errorf("get with cache: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return out, fmt.Errorf("json unmarshal: %w", err)
	}
	return out, nil
}

func (p *Pokemon) getName() string {
	return p.Name
}

func (p *Pokemon) getHeight() int {
	return p.Height
}

func (p *Pokemon) getWeight() int {
	return p.Weight
}

func (p *Pokemon) getStats() map[string]int {
	stats := map[string]int{}
	for _, stat := range p.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}
	return stats
}

func (p *Pokemon) getTypes() []string {
	var types []string
	for _, pokemonType := range p.Types {
		types = append(types, pokemonType.Type.Name)
	}
	return types
}
