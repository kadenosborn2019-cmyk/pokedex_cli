package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config) error {
	var areaList locationAreaList
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := http.Get(*cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
