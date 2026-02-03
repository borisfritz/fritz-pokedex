package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getResponse(url string) ([]byte, error) {
	log.Printf("Attempting ")
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get(%v) failed: %w", url, err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\n body: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(resp.Body) failed: %w", err)
	}
	return body, nil
}
