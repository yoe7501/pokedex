package main

import "fmt"

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex: ")
	if len(cfg.pokedex) == 0 {
		fmt.Println("..... is Empty")
	}
	for k := range cfg.pokedex {
		fmt.Printf("- %s\n", k)
	}

	return nil
}
