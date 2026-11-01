package main

func main() {
	locationAPI := "https://pokeapi.co/api/v2/location-area/"
	cfg := &config{
		next:     &locationAPI,
		prev:     nil,
		commands: getCommands(),
	}
	startRepl(cfg)
}
