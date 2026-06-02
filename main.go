package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	page        *pagination
}

type LocationAreaResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type pagination struct {
	next     string
	previous string
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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays the next 20 places
mapb: Displays the previous 20 places`)
	return nil
}

func newCommandMap(page *pagination) func() error {
	return func() error {
		res, err := http.Get(page.next)
		if err != nil {
			return fmt.Errorf("map: http get error: %w", err)
		}
		defer res.Body.Close()
		var out LocationAreaResponse
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&out); err != nil {
			return fmt.Errorf("map: body decoder error: %w", err)
		}
		page.previous = out.Previous
		page.next = out.Next
		for _, locationArea := range out.Results {
			fmt.Println(locationArea.Name)
		}
		return nil
	}
}

func newCommandMapb(page *pagination) func() error {
	return func() error {
		if page.previous == "" {
			fmt.Println("you're on the first page")
			return nil
		}
		res, err := http.Get(page.previous)
		if err != nil {
			return fmt.Errorf("mapb: http get error: %w", err)
		}
		defer res.Body.Close()
		var out LocationAreaResponse
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&out); err != nil {
			return fmt.Errorf("mapb: body decoder error: %w", err)
		}
		page.previous = out.Previous
		page.next = out.Next
		for _, locationArea := range out.Results {
			fmt.Println(locationArea.Name)
		}
		return nil
	}
}

func main() {
	mapPage := &pagination{next: "https://pokeapi.co/api/v2/location-area?limit=20&offset=0", previous: ""}
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			page:        &pagination{},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			page:        &pagination{},
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    newCommandMap(mapPage),
			page:        mapPage,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    newCommandMapb(mapPage),
			page:        mapPage,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf(`scanner error: %v`, err)
			break
		}
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]
		cmd, ok := supportedCommands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Printf("%s: callback error: %v\n", cmd.name, err)
		}
	}
}
