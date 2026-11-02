package main

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "list help commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "list 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list pokemon in city explore <city-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon catch <pokemon name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect pokemon if already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list pokemon in pokedex",
			callback:    commandPokedex,
		},
	}
	return commands
}
