package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cleanInput(" hi there!"))
}

// Split the users input into words and sanitize
func cleanInput(text string) []string {
	words := strings.Split(strings.TrimSpace(text), " ")

	return words
}
