go-linegraph
===

Generate line graph.

```golang
package main

import (
 "fmt"
 "math/rand"
 "sort"

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


 value := []float64{}

 d := pickup(1, 100, 21)

 for i := range d {
  j := rand.Intn(i + 1)
  d[i], d[j] = d[j], d[i]
 }

 for _, i := range d {
  value = append(value, float64(i))
 }

 graph.Data = value

 fmt.Println(graph.Data)
 fmt.Println(graph.Braille())
 fmt.Println(graph.Sparkline())

}
```

```
21 [83 49 59 63 32 22 81 27 61 64 90 71 48 91 43 3 54 16 65 17 92]
21 [⣧  ⣶  ⣄  ⣧  ⣶  ⣷  ⣼  ⡄  ⣆  ⣆  ⡇]
21 [▃ ▂ ▂ ▂ ▂ ▁ ▃ ▂ ▂ ▂ ▃ ▃ ▂ ▃ ▂ ▁ ▂ ▁ ▂ ▁ ▃]
```
