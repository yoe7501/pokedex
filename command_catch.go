package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatch(cfg *config) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + *cfg.pokemon
	var animal Pokemon
	res, err := http.Get(url)
	if err != nil {
		errors.New("Trouble retrieving pokemon")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		errors.New("issue reading res.body")
	}
	err = json.Unmarshal(data, &animal)
	if err != nil {
		errors.New("issue unmarshalling pokemon")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", animal.Name)
	if rand.Intn(animal.BaseExperience) > 40 {
		fmt.Printf("%s escaped!\n", animal.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", animal.Name)
	cfg.pokedex[animal.Name] = animal
	return nil
}
