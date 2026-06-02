package main

import "net/http"

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type configuration struct {
	Client   *http.Client
	Next     string
	Previous string
}

type LocationArea struct {
	id                   int
	name                 string
	gameIndex            int
	encounterMethodRates []EncounderMethodRate
	location             NamedAPIResource
	names                []Name
	pokemonEncounters    []PokemonEncounter
}

type EncounderMethodRate struct {
	encounterMethod NamedAPIResource
	versionDetails  []EncounterVersionDetail
}

type EncounterVersionDetail struct {
	rate    int
	version NamedAPIResource
}

type PokemonEncounter struct {
	pokemon        NamedAPIResource
	versionDetails []VersionEncounterDetail
}

type Name struct {
	name     string
	language NamedAPIResource
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionEncounterDetail struct {
	version          NamedAPIResource
	maxChance        int
	encounterDetails []Encounter
}

type Encounter struct {
	minLevel        int
	maxLevel        int
	conditionValues []NamedAPIResource
	chance          int
	method          NamedAPIResource
}

func main() {
	pokeClient := http.DefaultClient
	cfg := &configuration{Client: pokeClient, Next: "https://pokeapi.co/api/v2/location-area?limit=20&offset=0", Previous: ""}
	startRepl(cfg)
}
