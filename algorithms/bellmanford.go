package algorithms

import (
	"fmt"
	graph "github.com/colinwilcox1967/golanggraphs/graph"
)

const (
	SOME_RELATIVELY_BIG_FLOAT_VALUE = 999999.99
)

// Calculates the shortest distance from a vertex to all other vertices in a weighted graph
func BellmanFordAlgorithm(g *graph.Graph, sourceNodeId uint64) ([]float64, []float64) {
	var distance, predecesor []float64

	nodes := g.GetNodeList()
	arcs := g.GetArcList()

	distance = make([]float64, len(nodes))
	predecesor = make([]float64, len(nodes))

	// Setup the target arrays
	for index := 0; index < len(nodes); index++ {
		// infinitly far away and no previous noe
		distance[index] = SOME_RELATIVELY_BIG_FLOAT_VALUE
		predecesor[index] = -1
	}

	// trivial case
	sourceNodeIndex := g.FindNodeIndexWithId (sourceNodeId)
	distance[sourceNodeIndex] = 0.0

	// Repeatedly relax edges, considering the trivial case of a single node graph
	for index := 1; index <= len(nodes)-1; index++ {
		for _, arc := range arcs {
			from := arc.FromNodeId(g)
			fromIndex := g.FindNodeIndexWithId(from)
			to := arc.ToNodeId(g)
			toIndex := g.FindNodeIndexWithId(to)
			weight := arc.Weight()

			if distance[toIndex] > distance[fromIndex]+weight {
				distance[toIndex] = distance[fromIndex] + weight
				predecesor[toIndex] = float64(fromIndex)
			}
		}
	}

	// Check for negatively weighted cycles
	for _, arc := range arcs {
		from := arc.FromNodeId(g)
		fromIndex := g.FindNodeIndexWithId(from)

		to := arc.ToNodeId(g)
		toIndex := g.FindNodeIndexWithId(to)
		weight := arc.Weight()
		
		if distance[toIndex] > distance[fromIndex]+weight {
			fmt.Printf("The graph contains a negative-weight cycle\n")
		}
	}

	return distance, predecesor
}