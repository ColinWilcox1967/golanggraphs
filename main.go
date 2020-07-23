package main

import (
	"fmt"
	fileutilities "github.com/colinwilcox1967/golangfileutilities"
	graph "github.com/colinwilcox1967/golanggraphs/graph"
	"math"
	"strconv"
	"strings"
)

const (
	MAXUINT = ^uint(0)
	MAXINT  = int(MAXUINT >> 1)
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

func loadGraphDefinitionFromFile(g graph.Graph, filename string) error {

	//N:<id>-<value>
	//A:<from>-<to>-<weight>

	var nodeDetails []string
	var arcDetails []string
	var err error
	var lines []string

	if err, lines = fileutilities.ReadFileAsLines(filename); err == nil {

		// preprocess all the lines before parsing
		for index, _ := range lines {
			lines[index] = trimStringLeftAndRight (lines[index])	
		}

		// scan all nodes the scan all arcs
		for _, line := range lines {
			line = strings.ToUpper(line)
			if len(line) > 2 && line[0:2] == "N:" {
				newNodes := strings.Split(line[2:], ",")

				for _, node := range newNodes {
					nodeDetails = append(nodeDetails, node)
				}
			}
		}

		// add in any arc definitions
		for _, line := range lines {
			line = strings.ToUpper(line)
			if len(line) > 2 && line[0:2] == "A:" {
				newArcs := strings.Split(line[2:], ",")
				for _, arc := range newArcs {
					arcDetails = append(arcDetails, arc)
				}
			}
		}

		// now build the graph will the nodes then connect them up
		for _, node := range nodeDetails {
			// each node item is of form <id>-<value>
			index := strings.Index(node, "-")

			id, _ := strconv.ParseInt(node[:index], 10, 64)
			value, _ := strconv.ParseFloat(node[index+1:], 64)

			g.AddNode(uint64(id), float64(value))
		}

		for _, arc := range arcDetails {
			firstDash := strings.Index(arc, "-")
			secondDash := strings.Index(arc[firstDash+1:], "-")
			from, _ := strconv.Atoi(arc[:firstDash])
			to, _ := strconv.Atoi(arc[firstDash+1 : secondDash])
			weight, _ := strconv.ParseFloat(arc[secondDash+1:], 64)

			g.AddArc(uint64(from), uint64(to), float64(weight))
		}

		return nil
	}

	return err
}

// Calculates the shortest distance from a vertex to all other vertices in a weighted graph
func BellmanFordAlgorithm(g *graph.Graph, sourceNodeId uint64) ([]float64, []float64) {
	var distance, predecesor []float64

	nodes := g.GetNodeList()
	arcs := g.GetArcList()

	distance = make([]float64, len(nodes))
	predecesor = make([]float64, len(nodes))

	// Setup the target arrays
	for index, _ := range nodes {
		// infinitly far away and no previous noe
		distance[index] = math.MaxFloat64
		predecesor[index] = -1
	}

	// trivial case
	distance[sourceNodeId] = 0.0

	// Repeatedly relax edges, considering the trivial case of a single node graph
	for index := 1; index < len(nodes)-1; index++ {
		for _, arc := range arcs {
			from := arc.FromNodeId(g)
			to := arc.ToNodeId(g)
			weight := arc.Weight()
			if distance[from]+weight < distance[to] {
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
		if distance[from]+weight < distance[to] {
			fmt.Printf("The graph contains a negative-weight cycle\n")
		}
	}

	return distance, predecesor
}

func showBellmanFordAlgorithm(g *graph.Graph) {
	
	
	//	if err := loadGraphDefinitionFromFile ("BELLMANFORD.TXT"); err == nil {
	//		distance, predecessor := BellmanFordAlgorithm (g,0)
	//
	//		// print out the results here
	//	}

	distance, predecessor := BellmanFordAlgorithm (g,0)

	for index := 0; index < len(g.GetNodeList()); index++ {
		fmt.Printf ("[%02d] %f %f\n", index, distance[index], predecessor[index])
	}

}

func main() {
	g := graph.GetNewGraphInstance()

	g.AddNode(0, 10)
	g.AddNode(1, 5)
	g.AddArc(0, 1, 3)

	showBellmanFordAlgorithm(&g)
}
