package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
)

func AttemptCapture(cfg *Configuration, pokemon string) (err error) {
	captureRate, err := GetCaptureRate(cfg, pokemon)
	if err != nil {
		return fmt.Errorf("get capture rate: %w", err)
	}
	captureAttempt := rand.Int31n(256)
	if captureAttempt > captureRate {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.Pokedex.Add(cfg, pokemon)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon)
	return nil
}

func GetCaptureRate(cfg *Configuration, pokemon string) (captureRate int32, err error) {
	species, err := getPokemonSpecies(cfg, pokemon)
	if err != nil {
		return 0, fmt.Errorf("get pokemon species: %w", err)
	}
	return int32(species.CaptureRate), nil
}

func getPokemonSpecies(cfg *Configuration, pokemon string) (out PokemonSpecies, err error) {
	speciesURL, err := url.JoinPath(BaseURL, "/pokemon-species/", pokemon)
	if err != nil {
		return out, fmt.Errorf("url join path: %w", err)
	}
	data, err := getWithCache(cfg, speciesURL)
	if err != nil {
		return out, fmt.Errorf("get with cache: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return out, fmt.Errorf("json unmarshal: %w", err)
	}
	return out, nil
}
