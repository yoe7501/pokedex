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
		cmd.callback(cfg)
	}
}
