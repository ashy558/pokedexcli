package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func PrintLocationAreaResponse(url string, cfg *poketypes.Configuration) error {
	data, ok := cfg.Cache.Get(url)
	if !ok {
		fmt.Println("cache miss!")
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("http get error: %w", err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("body read error: %w", err)
		}
		cfg.Cache.Add(url, data)
	} else {
		fmt.Println("cache hit!")
	}
	var out LocationAreaResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("body unmarshal error: %w", err)
	}
	for _, locationArea := range out.Results {
		fmt.Println(locationArea.Name)
	}
	cfg.Previous = out.Previous
	cfg.Next = out.Next
	return nil
}
