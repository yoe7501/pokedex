package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type locationDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config) error {
	var detail locationDetail
	url := "https://pokeapi.co/api/v2/location-area/" + *cfg.city
	if value, ok := cfg.cache.Get(url); ok {
		err := json.Unmarshal(value, &detail)
		if err != nil {
			return errors.New("error unmarshalling data")
		}
	} else {
		res, err := http.Get(url)
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.New("issuue reading response body")
		}
		err = json.Unmarshal(data, &detail)
		if err != nil {
			return errors.New("Issue unmarshalling data")
		}

		cfg.cache.Add(url, data)
	}
	for _, value := range detail.PokemonEncounters {
		fmt.Println(value.Pokemon.Name)
	}
	return nil
}
