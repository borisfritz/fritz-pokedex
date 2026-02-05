package pokeapi

import (
	"fmt"
)

func (c *Client) GetLocationAreaBatch(url string) (LocationAreaBatch, error) {
	return decodeConfig[LocationAreaBatch](c, url)
}

func (c *Client) GetLocationAreaEndpoint(url string) (LocationAreaEndpoint, error) {
	return decodeConfig[LocationAreaEndpoint](c, url)
}

func (l LocationAreaBatch) PrintNames() {
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
}

func (l LocationAreaEndpoint) PrintPokemon(){
	for _, encounter := range l.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}
