package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *configuration) error {
	res, err := http.Get(cfg.Next)
	if err != nil {
		return fmt.Errorf("map: http get error: %w", err)
	}
	defer res.Body.Close()
	var out LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&out); err != nil {
		return fmt.Errorf("map: body decoder error: %w", err)
	}
	cfg.Previous = out.Previous
	cfg.Next = out.Next
	for _, locationArea := range out.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
