package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct {
	name		string
	description	string
	callback func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Turns off the Pokedex",
			callback: callbackExit,
		},
	}
}

func startRepl() {
	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user
		fmt.Print(" >")
		
		// Read user prompt
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		// Empty prompt (mashing return)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.callback()

	}
}


// sanitize user input before eval
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
