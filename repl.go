package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex >")
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
			err := command.callback(cfg)
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
	commands map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
