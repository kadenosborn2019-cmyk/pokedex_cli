package main

func main() {
	apiURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := &config{
		nextLocationsURL: &apiURL,
		prevLocationsURL: &apiURL,
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
			"mapb": {
				name:        "mapb",
				description: "mapb is a map back function that displays the previous pagination of locations",
				callback:    commandMapb,
			},
		},
		locations: locationAreaList{},
	}

	startRepl(cfg)
}
