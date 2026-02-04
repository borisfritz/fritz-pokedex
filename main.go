package main

import (
	"time"
	"github.com/borisfritz/fritz-pokedex/internal/pokeapi"
)

type replConfig struct {
	Next *string
	Prev *string
	Client *pokeapi.Client
}

func main() {
	cfg := &replConfig{
		Client: pokeapi.NewClient(5 * time.Second, time.Minute * 5),
	}
	startRepl(cfg)
}
