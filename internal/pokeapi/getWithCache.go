package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ashy558/pokedexcli/internal/poketypes"
)

func getWithCache(cfg *poketypes.Configuration, url string) ([]byte, error) {
	data, ok := cfg.Cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return data, fmt.Errorf("http get error: %w", err)
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return data, fmt.Errorf("body read error: %w", err)
		}
		cfg.Cache.Add(url, data)
	}
	return data, nil
}
