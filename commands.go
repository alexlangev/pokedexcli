package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name			string
	description		string
	callback		func() error
}

// var commands = map[string]cliCommand {
// 	"exit": {
// 		name:			"exit",
// 		description: 	"exit the Pokedex",
// 		callback: 		commandExit,
// 	},
// 	"help": {
// 		name:			"help",
// 		description:	"Displays a help message",
// 		callback:		commandHelp,
// 	},
// }

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"exit": {
			name:			"exit",
			description: 	"exit the Pokedex",
			callback: 		commandExit,
		},
	} 	
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) // 0 means success by convention
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
