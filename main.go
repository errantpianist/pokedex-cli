package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) > 0 {
		fmt.Printf("Your command was: %s\n", words[0])
	} else {
		fmt.Println("Empty command entered.")
	}
	}
}
