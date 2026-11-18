package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config, s string) error {
	var areaList locationAreaList
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	if hit, ok := cfg.cache.Get(*cfg.nextLocationsURL); ok {
		err := json.Unmarshal(hit, &areaList)
		if err != nil {
			return err
		}
		for _, name := range areaList.Results {
			fmt.Println(name.Name)
		}

	}
	resp, err := http.Get(*cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("reader.Body.Close() error: %v\n", err)
		}
	}()
	err = json.NewDecoder(resp.Body).Decode(&areaList)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = areaList.Next
	cfg.prevLocationsURL = areaList.Previous

	for _, name := range areaList.Results {
		fmt.Println(name.Name)
	}

	return nil
}
