package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := http.Get(*cfg.prevLocationsURL)
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
