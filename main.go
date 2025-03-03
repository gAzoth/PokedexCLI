package main

import (
	"time"

	"github.com/gAzoth/PokedexCLI/internal/pokeapi"
)

func main() {
	pokeCLient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeCLient,
	}

	startRep1(cfg)
}
