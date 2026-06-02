package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
exit: Exit the Pokedex`)
	return nil
}

func main() {
	supportedCommands := map[string]cliCommand{
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
