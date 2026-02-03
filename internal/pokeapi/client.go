package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) getResponse(url string) ([]byte, error) {
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
