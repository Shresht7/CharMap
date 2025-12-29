package charmap

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

// Implements the fuzzy.Source interface
// SymbolList implements fuzzy.Source for fuzzy searching.
type SymbolList []Symbol

func (s SymbolList) Len() int {
	return len(s)
}

// Used for fuzzy search
func (s SymbolList) String(i int) string {
	symbol := s[i]
	return strings.Join(
		[]string{
			symbol.Symbol,
			strings.Join(symbol.Keywords, " "),
			symbol.Unicode,
			symbol.Decimal,
			symbol.Latex,
			symbol.Category,
			symbol.Description,
		},
		" ",
	)
}
