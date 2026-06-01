package main

import "strings"

func cleanInput(text string) []string {
	if len(strings.TrimSpace(text)) == 0 {
		return []string{}
	}
	lowerText := strings.ToLower(text)
	cleanWords := strings.Fields(lowerText)
	return cleanWords
}
