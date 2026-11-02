package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	value, ok := cfg.pokedex[*cfg.pokemon]
	if !ok {
		return errors.New("Pokemon not in pokedex (pokemon has not been caught")
	}
	fmt.Println("Name: " + value.Name)
	fmt.Printf("Base Experience: %v\n", value.BaseExperience)
	fmt.Printf("Height: %v\n", value.Height)
	fmt.Printf("Weight: %v\n", value.Weight)
	fmt.Println("Stats:")
	for _, s := range value.Stats {
		fmt.Printf("- %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range value.Types {
		fmt.Println("- " + t.Type.Name)
	}
	return nil
}
