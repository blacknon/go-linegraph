package linegraph

import (
	"fmt"
)

var (
	graphSymbols = map[string]string{
		"00": "⠀", "01": "⢀", "02": "⢠", "03": "⢰", "04": "⢸",
		"10": "⡀", "11": "⣀", "12": "⣠", "13": "⣰", "14": "⣸",
		"20": "⡄", "21": "⣄", "22": "⣤", "23": "⣴", "24": "⣼",
		"30": "⡆", "31": "⣆", "32": "⣦", "33": "⣶", "34": "⣾",
		"40": "⡇", "41": "⣇", "42": "⣧", "43": "⣷", "44": "⣿",
	}
)

func usageToSymbol(current, next float64) string {
	getSymbolIndex := func(value float64) int {
		if value < 10 {
			return 0
		} else if value < 25 {
			return 1
		} else if value < 50 {
			return 2
		} else if value < 75 {
			return 3
		} else {
			return 4
		}
	}

	currentIndex := getSymbolIndex(current)
	nextIndex := getSymbolIndex(next)
	key := fmt.Sprintf("%d%d", currentIndex, nextIndex)
	return graphSymbols[key]
}

func (g *Graph) Braille() (result []string) {
	width := len(g.Data)
	if width == 0 {
		return
	}

	graphArray := make([]string, width)
	for i := 0; i < width; i += 2 {
		current := g.Data[i]
		next := 0.0
		if i < width-1 {
			next = g.Data[i+1]
		}

		symbol := usageToSymbol(current, next)
		graphArray[i] = symbol

	}

	for _, symbol := range graphArray {
		result = append(result, symbol)
	}

	return
}
