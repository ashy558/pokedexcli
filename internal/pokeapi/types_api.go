package pokeapi

type listLocationAreasResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             NamedAPIResource      `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []EncounterVersionDetail `json:"version_details"`
}

type EncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
}

type EncounterDetail struct {
	Chance          int              `json:"chance"`
	ConditionValues []any            `json:"condition_values"`
	MaxLevel        int              `json:"max_level"`
	Method          NamedAPIResource `json:"method"`
	MinLevel        int              `json:"min_level"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}
