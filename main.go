package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt string = "Pokedex > "

type appState struct {
	Next string
	Previous string
}

func main() {
	var state appState 
	state.Next = "https://pokeapi.co/api/v2/location-area/"
	state.Previous = ""
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := cleanInput(scanner.Text())

		// is a valid command?
		cmd, ok := getCommands()[input[0]]
		if ok {
			err := cmd.callback(&state)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Split the users input into words and sanitize
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))

	return words
}
