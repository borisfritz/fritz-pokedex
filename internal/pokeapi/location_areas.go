package pokeapi

import (
	"fmt"
)

//NOTE: LOCATION_AREA API structs and client methods at 'https://pokeapi.co/api/v2/'

//NOTE: Batch LOCATION-AREA data for pagination at 'https://pokeapi.co/api/v2/location-area/'
type LocationAreaBatch struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
func (l LocationAreaBatch) PrintNames() {
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
}

func (c *Client) GetLocationAreaBatch(url string) (LocationAreaBatch, error) {
	return decodeConfig[LocationAreaBatch](c, url)
}

//NOTE: Specific LOCATION_AREA data for individual location areas at 'https://pokeapi.co/api/v2/location-area/{id or name}'
type LocationAreaEndpoint struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
func (l LocationAreaEndpoint) PrintPokemon(){
	for _, encounter := range l.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}

func (c *Client) GetLocationAreaEndpoint(url string) (LocationAreaEndpoint, error) {
	return decodeConfig[LocationAreaEndpoint](c, url)
}
