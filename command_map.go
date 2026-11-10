package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	resp, err := http.Get("http://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cfg.locations)
	if err != nil {
		return err
	}
	for _, name := range cfg.locations.Results {
		fmt.Println(name.Name)
	}
	return nil
}
