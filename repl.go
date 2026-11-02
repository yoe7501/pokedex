package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yoe7501/pokedex/internal/pokecache"
)

type config struct {
	next     *string
	prev     *string
	commands map[string]cliCommand
	cache    *pokecache.Cache
	city     *string
	pokemon  *string
	pokedex  map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			fmt.Println("Must enter some kind of command type help for more information")
			continue
		}
		cmd, ok := cfg.commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if cmd.name == "explore" {
			if len(words) != 2 {
				fmt.Println("Explore must be used with a city name ")
				continue
			}
			cfg.city = &words[1]
		}
		if cmd.name == "catch" {
			if len(words) != 2 {
				fmt.Println("Catch must be used with a pokemon name")
				continue
			}
			cfg.pokemon = &words[1]
		}
		if cmd.name == "inspect" {
			if len(words) != 2 {
				fmt.Println("Inspect must be used with a pokemon name")
				continue
			}
			cfg.pokemon = &words[1]
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
