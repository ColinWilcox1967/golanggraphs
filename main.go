package main

import (
	"fmt"
	fileutilities "github.com/colinwilcox1967/golangfileutilities"
	graph "github.com/colinwilcox1967/golanggraphs/graph"
	"strconv"
	"strings"
)

var (
	g graph.Graph
)

func loadGraphDefinitionFromFile(filename string) error {

	//N:<id>-<value>
	//A:<from>-<to>-<weight>

	var nodeDetails []string
	var arcDetails []string
	var err error
	var lines []string

	if err, lines = fileutilities.ReadFileAsLines(filename); err == nil {

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

func main() {

	g = graph.CreateNewGraph()
	g.AddNode(1, 10)
	g.AddNode(2, 5)
	g.AddNode(3, 7)
	g.AddArc(1, 2, .1)
	g.AddArc(2, 3, .1)
	g.AddArc(3, 2, 0.1)
	g.AddArc(3, 1, .2)

	fmt.Printf("Number of nodes = %d\n", g.NodeCount())
	fmt.Printf("Number of arcs = %d\n", g.ArcCount())

	g.RemoveNodeById(2)

	//	if err := loadGraphDefinitionFromFile("test.txt"); err == nil {
	fmt.Printf("Number of nodes = %d\n", g.NodeCount())
	fmt.Printf("Number of arcs = %d\n", g.ArcCount())

}
