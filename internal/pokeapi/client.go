package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/borisfritz/fritz-pokedex/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) getResponse(url string) ([]byte, error) {
	//NOTE: Check if url is cached
	if val, ok := c.cache.Get(url); ok {
		return val, nil
	}
	
	//NOTE: If not, 'GET' request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest('GET', %v, nil) failed: %w", url, err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Get(%v) failed: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\n body: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(resp.Body) failed: %w", err)
	}
	return body, nil
}

func decodeConfig[T any](c *Client, url string) (T, error) {
	var data T
	body, err := c.getResponse(url)
	if err != nil {
		return data, fmt.Errorf("getResponse(%v) failed: %w", url, err)
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data, fmt.Errorf("json.Unmarshal(body, &data) failed: %w", err)
	}
	c.cache.Add(url, body)
	return data, nil
}
