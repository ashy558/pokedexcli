package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Printf(`an error occurred while scanning: %v`, err)
			break
		}
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}
