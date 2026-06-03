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

type Ability struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type ItemVersionDetails struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIResource `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIResource     `json:"item"`
	VersionDetails []ItemVersionDetails `json:"version_details"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type Move struct {
	Move                NamedAPIResource      `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type OtherSprites struct {
	DreamWorld      DreamWorldSprite      `json:"dream_world"`
	Home            HomeSprite            `json:"home"`
	OfficialArtwork OfficialArtworkSprite `json:"official-artwork"`
	Showdown        ShowdownSprite        `json:"showdown"`
}

type DreamWorldSprite struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  string `json:"front_female"`
}

type HomeSprite struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type OfficialArtworkSprite struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type ShowdownSprite struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type VersionsSprites struct {
	GenerationI struct {
		RedBlue struct {
			BackDefault  string `json:"back_default"`
			BackGray     string `json:"back_gray"`
			FrontDefault string `json:"front_default"`
			FrontGray    string `json:"front_gray"`
		} `json:"red-blue"`
		Yellow struct {
			BackDefault  string `json:"back_default"`
			BackGray     string `json:"back_gray"`
			FrontDefault string `json:"front_default"`
			FrontGray    string `json:"front_gray"`
		} `json:"yellow"`
	} `json:"generation-i"`
	GenerationIi struct {
		Crystal struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"crystal"`
		Gold struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"gold"`
		Silver struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"silver"`
	} `json:"generation-ii"`
	GenerationIii struct {
		Emerald struct {
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"emerald"`
		FireredLeafgreen struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"firered-leafgreen"`
		RubySapphire struct {
			BackDefault  string `json:"back_default"`
			BackShiny    string `json:"back_shiny"`
			FrontDefault string `json:"front_default"`
			FrontShiny   string `json:"front_shiny"`
		} `json:"ruby-sapphire"`
	} `json:"generation-iii"`
	GenerationIv struct {
		DiamondPearl struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"diamond-pearl"`
		HeartgoldSoulsilver struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"heartgold-soulsilver"`
		Platinum struct {
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"platinum"`
	} `json:"generation-iv"`
	GenerationV struct {
		BlackWhite struct {
			Animated struct {
				BackDefault      string `json:"back_default"`
				BackFemale       string `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  string `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"animated"`
			BackDefault      string `json:"back_default"`
			BackFemale       string `json:"back_female"`
			BackShiny        string `json:"back_shiny"`
			BackShinyFemale  string `json:"back_shiny_female"`
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"black-white"`
	} `json:"generation-v"`
	GenerationVi struct {
		OmegarubyAlphasapphire struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"omegaruby-alphasapphire"`
		XY struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"x-y"`
	} `json:"generation-vi"`
	GenerationVii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  string `json:"front_female"`
		} `json:"icons"`
		UltraSunUltraMoon struct {
			FrontDefault     string `json:"front_default"`
			FrontFemale      string `json:"front_female"`
			FrontShiny       string `json:"front_shiny"`
			FrontShinyFemale string `json:"front_shiny_female"`
		} `json:"ultra-sun-ultra-moon"`
	} `json:"generation-vii"`
	GenerationViii struct {
		Icons struct {
			FrontDefault string `json:"front_default"`
			FrontFemale  string `json:"front_female"`
		} `json:"icons"`
	} `json:"generation-viii"`
}

type Sprites struct {
	BackDefault      string          `json:"back_default"`
	BackFemale       string          `json:"back_female"`
	BackShiny        string          `json:"back_shiny"`
	BackShinyFemale  string          `json:"back_shiny_female"`
	FrontDefault     string          `json:"front_default"`
	FrontFemale      string          `json:"front_female"`
	FrontShiny       string          `json:"front_shiny"`
	FrontShinyFemale string          `json:"front_shiny_female"`
	Other            OtherSprites    `json:"other"`
	Versions         VersionsSprites `json:"versions"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stat struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIResource `json:"stat"`
}

type Type struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PastType struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []Type           `json:"types"`
}

type PastAbility struct {
	Generation NamedAPIResource `json:"generation"`
	Abilities  []Ability        `json:"abilities"`
}

type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []Ability          `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndices            []GameIndex        `json:"game_indices"`
	HeldItems              []HeldItem         `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []Move             `json:"moves"`
	Species                NamedAPIResource   `json:"species"`
	Sprites                Sprites            `json:"sprites"`
	Cries                  Cries              `json:"cries"`
	Stats                  []Stat             `json:"stats"`
	Types                  []Type             `json:"types"`
	PastTypes              []PastType         `json:"past_types"`
	PastAbilities          []PastAbility      `json:"past_abilities"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
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
	IsDefault bool             `json:"is_default"`
	Pokemon   NamedAPIResource `json:"pokemon"`
}
