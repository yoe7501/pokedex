package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	if text == ""{
		return []string{}
	}
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			fmt.Println("Must enter some kind of command type help for more information")
			continue
		}
		cmd, ok := commands[words[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}
}
