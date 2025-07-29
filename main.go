package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/errantpianist/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
		name string
		description string
		callback func(*config) error
}

type config struct {
	nextURL *string
	previousURL *string
}


func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return []string{}
	}
	words := strings.Fields(trimmed)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	locations, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.previousURL != nil {
		url = *cfg.previousURL
	}

	locations, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 location areas",
			callback: commandMapb,
		},
	
	}
}
	

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
	}
}
}
