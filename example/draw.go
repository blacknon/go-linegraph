package main

import (
	"fmt"
	"math/rand"

	linegraph "github.com/blacknon/go-linegraph"
)

func allKeys(m map[int]bool) []int {
	i := 0
	result := make([]int, len(m))
	for key, _ := range m {
		result[i] = key
		i++
	}
	return result
}

func pickup(min int, max int, num int) []int {
	numRange := max - min

	selected := make(map[int]bool)

	for counter := 0; counter < num; {
		n := rand.Intn(numRange) + min
		if selected[n] == false {
			selected[n] = true
			counter++
		}
	}
	keys := allKeys(selected)

	return keys
}

func main() {
	graph := &linegraph.Graph{}

	// データを生成
	value := []float64{}

	d := pickup(1, 100, 21)

	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}

	for _, i := range d {
		value = append(value, float64(i))
	}

	// データをセット
	graph.Data = value

	// グラフを描画
	braille := graph.Braille()
	sparkline := graph.Sparkline()
	fmt.Println(len(graph.Data), graph.Data)
	fmt.Println(len(braille), braille)
	fmt.Println(len(sparkline), sparkline)
}
