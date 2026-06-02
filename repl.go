package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
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

func startRepl(cfg *configuration) {
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
		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Printf("repl: callback error: %s", err)
			continue
		}
	}
}
