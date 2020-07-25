package main

import (
	"fmt"
	algorithms "github.com/colinwilcox1967/golanggraphs/algorithms"
	graph "github.com/colinwilcox1967/golanggraphs/graph"
	"strings"
)

func showBellmanFordAlgorithm(g *graph.Graph, filename string) {

	if err := g.LoadGraphDefinitionFromFile(filename); err == nil {

		distance, predecessor := algorithms.BellmanFordAlgorithm(g, 0)

		// print out the distances here
		fmt.Println ("[Bellman Ford Algorithm]")
		fmt.Println("Index   Distance  Predecesor")
		for index := 0; index < len(g.GetNodeList()); index++ {
			fmt.Printf("[%02d]    %6f  %6f\n", index, distance[index], predecessor[index])
		}
	} else {
		fmt.Printf("Unable to load file '%s'\n", strings.ToUpper(filename))
	}
}

func main() {
	g := graph.GetNewGraphInstance()

	showBellmanFordAlgorithm(&g, "testfiles/bellmanford.txt")
}
