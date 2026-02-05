package pokeapi

import (
	"fmt"
	"math/rand"
)

func (c *Client) GetPokemonData(url string) (PokemonData, error) {
	return decodeConfig[PokemonData](c, url)
}

func (p PokemonData) pokemonCaptureRate() float64 {
	const (
		lowExp = 40.0
		highExp = 400.0
		maxProb = 90.0
		minProb = 10.0
	)
	baseExp := float64(p.BaseExperience)
	if baseExp <= lowExp {
		return maxProb
	}
	if baseExp >= highExp {
		return minProb
	}
	rate := maxProb - (baseExp - lowExp) * ((maxProb - minProb)/(highExp - lowExp))
	return rate
}

func (p PokemonData) AttemptCapture() bool {
	chance := p.pokemonCaptureRate()
	roll := rand.Float64() * 100
	return roll <= chance
}

//NOTE: TEMPORARY function for testing Get and data gathering!

func (p PokemonData) PrintPokemonInfo() {
	fmt.Printf("Pokemon Name: %v\n", p.Name)
	fmt.Printf("Pokemon ID: %v\n", p.ID)
	fmt.Printf("Base Experience: %v\n", p.BaseExperience)
}
