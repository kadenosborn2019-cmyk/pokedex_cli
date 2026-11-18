package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandExplore(cfg *config, s string) error {
	var exploreArea locationAreaList
	resp, err := http.Get(*cfg.nextLocationsURL + s)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("could not close respBody: %v", err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(&exploreArea); err != nil {
		return err
	}

	stream, err := json.Marshal(&exploreArea)
	if err != nil {
		return err
	}

	if err := cfg.cache.Add(*cfg.nextLocationsURL+s, stream); err != nil {
		return err
	}
	return nil
}
