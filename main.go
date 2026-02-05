package main

import (
	"time"
	"github.com/borisfritz/fritz-pokedex/internal/pokeapi"
)

type replConfig struct {
	Next    *string
	Prev    *string
	Client  *pokeapi.Client
	Pokedex map[string]pokeapi.PokemonData
}

func main() {
	cfg := &replConfig{
		Client: pokeapi.NewClient(5 * time.Second, time.Minute * 5),
		Pokedex: make(map[string]pokeapi.PokemonData),
	}
	startRepl(cfg)
}
