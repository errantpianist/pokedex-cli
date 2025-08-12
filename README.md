<div align="center">
   <h1>🐾 Pokedex CLI</h1>
   <p>A command-line Pokedex application built in Go that interacts with the PokeAPI to explore locations, catch Pokemon, and manage your collection.</p>
</div>

---

## ✨ Features

- **Explore Pokemon World**: Navigate through location areas using `map` and `mapb` commands
- **Catch Pokemon**: Use the `catch` command to add Pokemon to your collection with randomized catch mechanics
- **Inspect Pokemon**: View detailed information about caught Pokemon with the `inspect` command
- **Manage Collection**: Keep track of all caught Pokemon with the `pokedex` command
- **API Integration**: Fetches real Pokemon data from the PokeAPI
- **Smart Caching**: Built-in caching system for improved performance and reduced API calls
- **Interactive REPL**: Command-line interface with persistent session

## 🚀 Getting Started

### Prerequisites

- Go 1.19+
- Git

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/errantpianist/pokedexcli.git
   cd pokedexcli
   ```

2. Build the application:

   ```sh
   go build
   ```

3. Run the Pokedex CLI:

   ```sh
   ./pokedexcli
   ```

### Usage

```sh
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays the next 20 location areas
mapb: Displays the previous 20 location areas
explore: Explore a location area for Pokemon
catch: Catch a Pokemon
inspect: View details of a caught Pokemon
pokedex: View all caught Pokemon
```

### Example Workflow

```sh
Pokedex > map
[shows location areas]

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - magikarp
 - wingull

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!
You may now inspect it with the inspect command.

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 45
Stats:
  -hp: 40
  -attack: 40
  -defense: 35
  -special-attack: 50
  -special-defense: 100
  -speed: 70
Types:
  - water
  - poison

Pokedex > pokedex
Your Pokedex:
 - tentacool
```

## 🏗️ Architecture

Command Pattern: Each command is implemented as a function with a consistent signature
Caching Layer: In-memory cache with automatic cleanup to reduce API calls
Config Management: Centralized configuration for API clients and user data
Error Handling: Comprehensive error handling for API calls and user input

## 🙏 Credits

Developed by errantpianist
Powered by PokeAPI
Inspired by the Boot.dev backend course

Gotta catch 'em all! 🌟
