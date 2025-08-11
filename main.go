package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/errantpianist/pokedexcli/internal/pokeapi"
	"github.com/errantpianist/pokedexcli/internal/pokecache"
)

type cliCommand struct {
		name string
		description string
		callback func(*config, ...string) error
}

type Pokemon struct {
	ID	 int    `json:"id"`
	Name string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats map[string]int `json:"stats"`
	Types []string `json:"types"`
}

type config struct {
	nextURL *string
	previousURL *string
	cache *pokecache.Cache
	pokedex map[string]Pokemon
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

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextURL != nil {
		url = *cfg.nextURL
	}

	locations, err := pokeapi.GetLocationAreas(url, cfg.cache)
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousURL == nil {
		fmt.Println("You are already at the first page of location areas.")
		return nil
	}

	locations, err := pokeapi.GetLocationAreas(*cfg.previousURL, cfg.cache)
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location area name")
	}

	areaName := args[0]
	area, err := pokeapi.GetLocationArea(areaName, cfg.cache)
	if err != nil {
		return fmt.Errorf("error fetching location area %s: %v", areaName, err)
	}

	fmt.Printf("Exploring %s:\n", area.Name)
	fmt.Println("Found Pokemon:")

	for _, encounter := range area.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a Pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := pokeapi.GetPokemon(pokemonName, cfg.cache)
	if err != nil {
		return fmt.Errorf("error catching Pokemon %s: %v", pokemonName, err)
	}

	rand.Seed(time.Now().UnixNano())
	catchChance := rand.Intn(pokemon.BaseExperience + 50)

	if catchChance < 50 {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect commend.")

		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]Pokemon)
		}

		stats := make(map[string]int)
		for _, stat := range pokemon.Stats {
			stats[stat.Stat.Name] = stat.BaseStat
		}

		types := make([]string, len(pokemon.Types))
		for i, t := range pokemon.Types {
			types[i] = t.Type.Name
		}


		cfg.pokedex[pokemonName] = Pokemon{
			ID: pokemon.ID,
			Name: pokemon.Name,
			Height: pokemon.Height,
			Weight: pokemon.Weight,
			Stats: stats,
			Types: types,
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
	
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a Pokemon name to inspect")
	}

	pokemonName := args[0]

	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	statNames := []string{"hp", "attack", "defense", "special-attack", "special-defense", "speed"}
	for _, statName := range statNames {
		if statValue, exists := pokemon.Stats[statName]; exists {
			fmt.Printf("  - %s: %d\n", statName, statValue)
		}
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}


	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.pokedex) == 0 {
		fmt.Println(" - (empty)")
		return nil
	}

	// Get names and ids and sort by id
	
	pokemonList := make([]string, 0, len(cfg.pokedex))

	for name, pokemon := range cfg.pokedex {
		pokemonList = append(pokemonList, fmt.Sprintf("%s (ID: %d)", name, pokemon.ID))
	}

	for _, name := range pokemonList {
		fmt.Println(" -", name)
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
		"explore": {
			name: "explore",
			description: "Explores a specific location area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catches a Pokemon by name",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "View details of a caught Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "View your caught Pokemon",
			callback: commandPokedex,
		},
	}
}
	

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Minute), // Cache for 5 minutes
		pokedex: make(map[string]Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
	}
}
}
