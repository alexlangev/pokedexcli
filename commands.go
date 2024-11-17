package main
import (
	"fmt"
	"os"
)

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func callbackExit() error{
	os.Exit(0) // terminate the program without errors
	return nil
}
