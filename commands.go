package main

func getCommands() map[string]cliCommand{
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name: "help",
			description: "list help commands",
			callback: commandHelp,
		},
	}
	return commands
}
