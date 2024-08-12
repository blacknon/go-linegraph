package linegraph

type Graph struct {
	// Data is the data to be graphed.
	Data []float64

	// Min and Max are the minimum and maximum values of the graph.
	Min float64
	Max float64
}
