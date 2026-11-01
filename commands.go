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
	}
	return commands
}
