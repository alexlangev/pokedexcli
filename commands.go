package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []location
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

	// Decode the JSON into a go struct
	decRes := locationResponse{}
	err = json.Unmarshal(body, &decRes)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range decRes.Results {
		fmt.Println(loc.Name)
	}

	// fmt.Println("DecRes.Results\n\n\n", decRes.Results[0].Name, decRes.Results[0].URL)

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
