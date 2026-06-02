package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ashy558/pokedexcli/internal/poketypes"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*poketypes.Configuration, string) error
}

func cleanInput(text string) []string {
	if len(strings.TrimSpace(text)) == 0 {
		return []string{}
	}
	lowerText := strings.ToLower(text)
	cleanWords := strings.Fields(lowerText)
	return cleanWords
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon present",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func startRepl(cfg *poketypes.Configuration) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf("repl: scanner error: %s\n", err)
			continue
		}
		cleanedInput := cleanInput(input)
		command := cleanedInput[0]
		args := ""
		if len(cleanedInput) > 1 {
			args = cleanedInput[1]
		}
		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, args); err != nil {
			fmt.Printf("repl: %s: %s\n", cmd.name, err)
			continue
		}
	}
}
