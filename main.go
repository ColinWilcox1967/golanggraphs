package main

import (
	"fmt"
	"strings"
	graph "github.com/colinwilcox1967/golanggraphs/graph"
)

const (
	SOME_RELATIVELY_BIG_FLOAT_VALUE = 999999.99
)



func trimStringLeftAndRight (line string) string {
	
	var leftIndex, rightIndex int
	
	leftIndex = 0
	for line[leftIndex] == ' ' && leftIndex < len(line) {
		leftIndex++ 	
	}
	newLine := line[:leftIndex]

	rightIndex = len(newLine)
	for newLine[rightIndex] == ' ' && rightIndex >= 0 {
		rightIndex--
	}
	newLine = newLine[rightIndex:]


	return newLine

}



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
	distance[sourceNodeId] = 0.0

	// Repeatedly relax edges, considering the trivial case of a single node graph
	for index := 1; index <= len(nodes)-1; index++ {
		for _, arc := range arcs {
			from := arc.FromNodeId(g)
			to := arc.ToNodeId(g)
			weight := arc.Weight()

			if distance[to] > distance[from]+weight {
				distance[to] = distance[from] + weight
				predecesor[to] = float64(from)
			}
		}
	}

	// Check for negatively weighted cycles
	for _, arc := range arcs {
		from := arc.FromNodeId(g)
		to := arc.ToNodeId(g)
		weight := arc.Weight()
		if distance[to] > distance[from]+weight {
			fmt.Printf("The graph contains a negative-weight cycle\n")
		}
	}

	return distance, predecesor
}

func showBellmanFordAlgorithm(g *graph.Graph, filename string) {
	
	if err := g.LoadGraphDefinitionFromFile (filename); err == nil {

		fmt.Println ("NODES")
		for index, node := range g.GetNodeList () {
			fmt.Printf ("[%02d] %d %d\n", index, node.Id(), node.Value ())
		}
		fmt.Println ("ARCS")
		for index, arc := range g.GetArcList () {
			fmt.Printf ("[%02d] %d %d %f\n", index, arc.FromNodeId(g), arc.ToNodeId (g), arc.Weight ())
		}


		distance, predecessor := BellmanFordAlgorithm (g,0)

		fmt.Println ("Outside")
		
		// print out the distances here
		for index := 0; index < len(g.GetNodeList()); index++ {
			fmt.Printf ("[%02d] %f %f\n", index, distance[index], predecessor[index])
		}
	} else {
		fmt.Printf ("Unable to load file '%s'\n", strings.ToUpper(filename))
	}
}

func main() {
	g := graph.GetNewGraphInstance()

	showBellmanFordAlgorithm(&g,"bellmanford.txt")
}
