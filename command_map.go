package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	resp, err := http.Get(*cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cfg.locations)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = cfg.locations.Next
	cfg.prevLocationsURL = cfg.locations.Previous

	for _, name := range cfg.locations.Results {
		fmt.Println(name.Name)
	}

	return nil
}
