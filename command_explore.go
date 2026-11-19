package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters []struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

func commandExplore(cfg *config, s []string) error {
	var encounters locationAreaList
	resp, err := http.Get(*cfg.nextLocationsURL + s[0] + "/")
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("could not close respBody: %v", err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(&encounters); err != nil {
		return err
	}

	stream, err := json.Marshal(&encounters)
	if err != nil {
		return err
	}

	if err := cfg.cache.Add(*cfg.nextLocationsURL+"/"+s[0], stream); err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", s[0])
	fmt.Println("Found Pokemon: ")
	for _, name := range encounters.PokemonEncounters {
		fmt.Printf("- %v\n", name.Pokemon.Name)
	}

	return nil
}
