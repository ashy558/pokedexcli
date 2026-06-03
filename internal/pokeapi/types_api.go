package pokeapi

type APIResource struct {
	URL string `json:"url"`
}

type EncounterDetail struct {
	Chance          int              `json:"chance"`
	ConditionValues []any            `json:"condition_values"`
	MaxLevel        int              `json:"max_level"`
	Method          NamedAPIResource `json:"method"`
	MinLevel        int              `json:"min_level"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type FlavorTextEntry struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedAPIResource `json:"language"`
	Version    NamedAPIResource `json:"version"`
}

type FormDescription struct {
	Description string           `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type Genus struct {
	Genus    string           `json:"genus"`
	Language NamedAPIResource `json:"language"`
}

type ListLocationAreasResponse struct {
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

type Name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokedexNumber struct {
	EntryNumber int              `json:"entry_number"`
	Pokedex     NamedAPIResource `json:"pokedex"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon                  `json:"pokemon"`
	VersionDetails []EncounterVersionDetail `json:"version_details"`
}

type PokemonSpecies struct {
	ID                   int                `json:"id"`
	Name                 string             `json:"name"`
	Order                int                `json:"order"`
	GenderRate           int                `json:"gender_rate"`
	CaptureRate          int                `json:"capture_rate"`
	BaseHappiness        int                `json:"base_happiness"`
	IsBaby               bool               `json:"is_baby"`
	IsLegendary          bool               `json:"is_legendary"`
	IsMythical           bool               `json:"is_mythical"`
	HatchCounter         int                `json:"hatch_counter"`
	HasGenderDifferences bool               `json:"has_gender_differences"`
	FormsSwitchable      bool               `json:"forms_switchable"`
	GrowthRate           NamedAPIResource   `json:"growth_rate"`
	PokedexNumbers       []PokedexNumber    `json:"pokedex_numbers"`
	EggGroups            []NamedAPIResource `json:"egg_groups"`
	Color                NamedAPIResource   `json:"color"`
	Shape                NamedAPIResource   `json:"shape"`
	EvolvesFromSpecies   NamedAPIResource   `json:"evolves_from_species"`
	EvolutionChain       APIResource        `json:"evolution_chain"`
	Habitat              NamedAPIResource   `json:"habitat"`
	Generation           NamedAPIResource   `json:"generation"`
	Names                []Name             `json:"names"`
	FlavorTextEntries    []FlavorTextEntry  `json:"flavor_text_entries"`
	FormDescriptions     []FormDescription  `json:"form_descriptions"`
	Genera               []Genus            `json:"genera"`
	Varieties            []Variety          `json:"varieties"`
}

type Variety struct {
	IsDefault bool    `json:"is_default"`
	Pokemon   Pokemon `json:"pokemon"`
}
