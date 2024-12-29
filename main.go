package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

const prompt string = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Println("Your command was:", input[0])
	}
}

// Split the users input into words and sanitize
func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))

	return words
}
