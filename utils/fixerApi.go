package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct to parse the JSON response from Fixer.io
type FixerResponse struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}

func GetConversionRate(apiKey, target string) (float64, error) {
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s&symbols=%s", apiKey, target)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var fixerResp FixerResponse
	if err := json.NewDecoder(resp.Body).Decode(&fixerResp); err != nil {
		return 0, err
	}

	if !fixerResp.Success {
		return 0, fmt.Errorf("failed to get exchange rate")
	}

	rate, exists := fixerResp.Rates[target]
	if !exists {
		return 0, fmt.Errorf("rate for %s not found", target)
	}

	return rate, nil
}
