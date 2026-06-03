package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func listLocationAreas(cfg *Configuration, url string) (out []string, err error) {
	data, err := getWithCache(cfg, url)
	if err != nil {
		return out, fmt.Errorf("getWithCache: %w", err)
	}
	var res ListLocationAreasResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return out, fmt.Errorf("body unmarshal error: %w", err)
	}
	cfg.Previous = res.Previous
	cfg.Next = res.Next
	for _, locationArea := range res.Results {
		out = append(out, locationArea.Name)
	}
	return out, nil
}

func getLocationArea(cfg *Configuration, area string) (out LocationArea, err error) {
	url, err := url.JoinPath(BaseURL, "/location-area/", area)
	if err != nil {
		return out, fmt.Errorf("URLJoinPath: %w", err)
	}
	data, err := getWithCache(cfg, url)
	if err != nil {
		return out, fmt.Errorf("getWithCache: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return out, fmt.Errorf("body unmarshal error: %w", err)
	}
	return out, nil
}

func PrintLocationAreas(cfg *Configuration, url string) error {
	areas, err := listLocationAreas(cfg, url)
	if err != nil {
		return fmt.Errorf("listLocationAreas: %w", err)
	}
	for _, area := range areas {
		fmt.Println(area)
	}
	return nil
}
