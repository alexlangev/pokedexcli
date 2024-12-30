package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"log"
	"encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 locations at a time",
			callback:    commandMap,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandMap() error {
	// make a GET to endpoint and handle errors
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}

	// "read" the body (encoded JSON) and handle errors
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locations := location{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", locations)

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) // 0 means success by convention
	return nil
}
