package main

import (
	"encoding/json"
	"os"
)

type Symbol struct {
	Symbol      string   `json:"symbol"`
	Category    string   `json:"category"`
	Unicode     string   `json:"unicode"`
	Decimal     string   `json:"decimal"`
	Latex       string   `json:"latex"`
	Keywords    []string `json:"keywords"`
	Description string   `json:"description"`
}

func LoadSymbols(path string) (map[string]Symbol, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var symbols map[string]Symbol
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&symbols); err != nil {
		return nil, err
	}

	return symbols, nil
}
