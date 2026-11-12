package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	var areaList locationAreaList
	resp, err := http.Get(*cfg.nextLocationsURL)
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
