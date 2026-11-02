package main

import (
	"time"

	"github.com/yoe7501/pokedex/internal/pokecache"
)

func main() {
	locationAPI := "https://pokeapi.co/api/v2/location-area/"
	c := pokecache.NewCache(5 * time.Minute)
	cfg := &config{
		next:     &locationAPI,
		prev:     nil,
		commands: getCommands(),
		cache:    c,
	}
	startRepl(cfg)
}
