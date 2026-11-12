package main

func main() {
	apiURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := config{
		commands: map[string]cliCommand{
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
				name:        "map",
				description: "Displays a pagination of locations from the pokeapi",
				callback:    commandMap,
			},
		},
		locations: locationAreaList{
			Next: &apiURL,
		},
	}

	startRepl(&cfg)
}
