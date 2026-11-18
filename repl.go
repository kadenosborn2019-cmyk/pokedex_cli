package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/kadenosborn2019-cmyk/pokedex_cli/internal/pokeCache"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		if scanner.Err() != nil {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := cfg.commands[commandName]
		if ok {
			err := command.callback(cfg, "")
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	clean := strings.Fields(lowerText)
	return clean
}

type config struct {
	commands         map[string]cliCommand
	nextLocationsURL *string
	prevLocationsURL *string
	cache            pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type NamedResources struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationAreaList struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []NamedResources `json:"results"`
}
