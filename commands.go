package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alexlangev/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*appState) error
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
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
			description: "Displays the name of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of the previous 20 locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandMap(state *appState, c *cacheEntry) error {
	// make a GET to endpoint and handle errors
	res, err := http.Get(state.Next)
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

	// add body to cache
	c.cache.Add(url, body)

	// Decode the JSON into a go struct
	decRes := locationResponse{}
	err = json.Unmarshal(body, &decRes)
	if err != nil {
		log.Fatal(err)
	}

	// update app state
	state.Next = decRes.Next
	state.Previous = decRes.Previous

	for _, loc := range decRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(state *appState) error {
	// make a GET to endpoint and handle errors
	res, err := http.Get(state.Previous)
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

	// update app state
	state.Next = decRes.Next
	state.Previous = decRes.Previous

	for _, loc := range decRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandHelp(state *appState) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func commandExit(state *appState) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0) // 0 means success by convention
	return nil
}
