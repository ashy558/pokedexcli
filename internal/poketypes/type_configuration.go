package poketypes

import (
	"net/http"

	"github.com/ashy558/pokedexcli/internal/pokecache"
)

type Configuration struct {
	Client   *http.Client
	Next     string
	Previous string
	Cache    *pokecache.Cache
}
