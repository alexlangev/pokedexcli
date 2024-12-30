package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt string = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := cleanInput(scanner.Text())

		// is a valid command?
		cmd, ok := getCommands()[input[0]]
		if ok {
			err := cmd.callback()
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
