package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationAreaListResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationAreaItem
}

type LocationAreaItem struct {
	Name string
	URL  string
}

func PrintLocationAreaData(url string) error {
	return nil
}

func GetLocationAreaData(url string) (LocationAreaListResponse, error) {
	body, err := getResponse(url)
	if err != nil {
		return LocationAreaListResponse{}, fmt.Errorf("getResponse(%v) failed: %w", url, err)
	}
	var data LocationAreaListResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreaListResponse{}, fmt.Errorf("json.Unmarshal(body, &data) failed: %w", err)
	}
	return data, nil
}
