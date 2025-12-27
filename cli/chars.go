package main

import (
	"encoding/json"
	"os"
	"strings"
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

func (s Symbol) String() string {
	return strings.Join(
		[]string{
			s.Symbol,
			s.Category,
			s.Description,
			strings.Join(s.Keywords, ", "),
			s.Unicode,
			s.Decimal,
			s.Latex,
		},
		" ",
	)
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

func GetAllSymbols(symbols map[string]Symbol) []Symbol {
	allSymbols := make([]Symbol, 0, len(symbols))
	for _, symbol := range symbols {
		allSymbols = append(allSymbols, symbol)
	}
	return allSymbols
}
