package linegraph

import (
	"math"
)

var (
	sparklineTicks = []string{
		"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█",
	}
)

func (g *Graph) Sparkline() (result []string) {

	if len(g.Data) == 0 {
		return
	}

	min := g.Min
	max := g.Max
	if min == 0 && max == 0 {
		min, max = g.Data[0], g.Data[0]
		for _, num := range g.Data {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
	}

	sizeRatio := float64(len(sparklineTicks)) / float64(len(g.Data))
	if sizeRatio > 1 {
		sizeRatio = 1
	}

	numTicks := int(sizeRatio * float64(len(sparklineTicks)))

	if numTicks < 1 {
		numTicks = 1
	}

	if math.Abs(max-min) < 0.0000001 {
		for range g.Data {
			result = append(result, sparklineTicks[0])
		}
	} else {
		scale := float64(numTicks-1) / (max - min)
		for _, n := range g.Data {
			tick := int((n-min)*scale + 0.5)
			if tick >= numTicks {
				tick = numTicks - 1
			}
			result = append(result, sparklineTicks[tick])
		}
	}

	return
}
