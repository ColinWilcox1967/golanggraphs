package graph

import (
"strings"
"fmt"
"strconv"
	fileutilities "github.com/colinwilcox1967/golangfileutilities"
)


const (
	KErrorNone               int = 0
	KErrorNodeNotFound       int = -1
	KErrorConflictingNodeIds int = -2
)

type Graph struct {
	nodes []Node
	arcs  []Arc
}

var (
gr Graph
)

func GetNewGraphInstance() Graph {

	gr.nodes = nil
	gr.arcs = nil

	return gr
}


//
// Graph methods
//
func (g * Graph)LoadGraphDefinitionFromFile(filename string) error {

	//N:<id>-<value>
	//A:<from>-<to>-<weight>

	var nodeDetails []string
	var arcDetails []string
	var err error
	var lines []string

	if err, lines = fileutilities.ReadFileAsLines(filename); err == nil {

		// preprocess all the lines before parsing
		for index := 0; index < len(lines);index++{
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

			if index >= 0 {
				id, _ := strconv.ParseInt(node[:index], 10, 64)
				value, _ := strconv.ParseFloat(node[index+1:], 64)

				g.AddNode(uint64(id), float64(value)) 
			} else {
				fmt.Printf ("Malformed node definition: '%s'. Skipping.\n", node)
			}

		}

		for _, arc := range arcDetails {

			fragments := strings.Split (arc, "-")

			
			if len(fragments) == 3 {
				from, _ := strconv.Atoi(fragments[0])
				to, _ := strconv.Atoi(fragments[1])
				weight, _ := strconv.ParseFloat(fragments[2], 64)

				g.AddArc(uint64(from), uint64(to), float64(weight))
			} else {
				fmt.Printf ("Malformed arc definition: '%s'. Skipping.\n", arc)
			}
		}
		return nil
	}


	return err
}
func (g *Graph) AddArc(fromNode, toNode uint64, weight float64) int {
	if g.uniqueId(fromNode) || g.uniqueId(toNode) {
		return KErrorConflictingNodeIds
	}

	var newArc Arc
	newArc.from = fromNode
	newArc.to = toNode
	newArc.weight = weight

	// add arc to graph arc list
	g.arcs = append(g.arcs, newArc)

	fromIndex := g.findNodeIndexWithId(fromNode)
	toIndex := g.findNodeIndexWithId(toNode)

	// hook up the nodes based on the arc information
	g.nodes[fromIndex].outbound = append(g.nodes[fromIndex].outbound, uint64(toNode))
	g.nodes[toIndex].inbound = append(g.nodes[toIndex].inbound, uint64(fromNode))

	return KErrorNone
}

func (g *Graph) AddNode(id uint64, value float64) int {
	if !g.uniqueId(id) {
		return KErrorNodeNotFound
	}

	// we have a new Id
	var newNode Node
	newNode.id = id
	newNode.value = value
	g.nodes = append(g.nodes, newNode)

	return KErrorNone
}


func (g *Graph)GetNodeList () []Node {
	return g.nodes
}

func (g *Graph)GetArcList () []Arc {
	return g.arcs
}

func (g *Graph) NodeInboundConnections(index int) []uint64 {
	return g.nodes[index].inbound
}

func (g *Graph) NodeOutboundConnections(index int) []uint64 {
	return g.nodes[index].outbound
}

func (g *Graph) IsConnected(fromNode, toNode uint64) bool {

	status, node := g.findNodeWithId(toNode)
	if status == KErrorNodeNotFound {
		return false
	}

	inboundArcs := node.InboundArcList()
	outboundArcs := node.OutboundArcList()

	for _, connectedNode := range inboundArcs {
		if connectedNode == fromNode {
			return true
		}
	}

	for _, connectedNode := range outboundArcs {
		if connectedNode == toNode {
			return true
		}
	}

	return false
}

func (g *Graph) NodeCount() int {

	return len(g.nodes)
}

func (g *Graph) ArcCount() int {
	return len(g.arcs)
}

func (g *Graph) OrphanedNodeIDs() []uint64 {
	var nodeIds []uint64

	for _, node := range g.nodes {
		if len(node.inbound) == 0 && len(node.outbound) == 0 {
			nodeIds = append(nodeIds, node.id)
		}
	}
	return nodeIds
}

func (g *Graph) RemoveNodeById(id uint64) bool {
	var index int

	if index = g.findNodeIndexWithId(id); index == KErrorNodeNotFound {
		return false // invalid node id in g
	}

	// Disconnect all inbound nodes
	for _, nodeId := range g.nodes[index].inbound {
		nodeReferenceIndex := g.findNodeIndexWithId(nodeId)

		// this needs to be streamlined later
		g.removeArcsBetweenTwoNodes(nodeId, id)

		// remove node reference
		if nodeReferenceIndex >= 0 {
			// remove arcs between these two nodes

			// remove node reference
			pos := findNodeIdInList(g.nodes[nodeReferenceIndex].outbound, id)
			if len(g.nodes[nodeReferenceIndex].outbound) == 1 {
				g.nodes[nodeReferenceIndex].outbound = nil
			} else {
				g.nodes[nodeReferenceIndex].outbound = append(g.nodes[nodeReferenceIndex].outbound[:pos], g.nodes[nodeReferenceIndex].outbound[pos+1:]...)
			}
		}
	}

	// disconnect all outbound nodex
	for _, nodeId := range g.nodes[index].outbound {
		nodeReferenceIndex := g.findNodeIndexWithId(nodeId)

		// this needs to be streamlined later
		g.removeArcsBetweenTwoNodes(id, nodeId)

		// remove node reference
		if nodeReferenceIndex >= 0 {
			pos := findNodeIdInList(g.nodes[nodeReferenceIndex].inbound, id)
			if len(g.nodes[nodeReferenceIndex].inbound) == 1 {
				g.nodes[nodeReferenceIndex].inbound = nil
			} else {
				g.nodes[nodeReferenceIndex].inbound = append(g.nodes[nodeReferenceIndex].inbound[:pos], g.nodes[index].inbound[pos+1:]...)
			}
		}
	}

	//remove the node itself
	g.nodes = append(g.nodes[:index], g.nodes[index+1:]...)

	return true
}

// private helper function

//
// just tidies up config file lines
//
func trimStringLeftAndRight (line string) string {
	
////	var leftIndex, rightIndex int
//	
//	leftIndex = 0
//	for line[leftIndex] == ' ' && leftIndex < len(line) {
//		leftIndex++ 	
//	}
//	newLine := line[leftIndex:]
//
//	rightIndex = len(newLine)-1
//	for newLine[rightIndex] == ' ' && rightIndex >= 0 {
//		rightIndex--
//	}
//	newLine = newLine[:rightIndex+1]


//	return newLine

return line

}

func (g *Graph) uniqueId(id uint64) bool {
	for _, node := range g.nodes {
		if node.Id() == id {
			return false
		}
	}
	return true
}

func (g *Graph) findNodeWithId(id uint64) (int, Node) {

	for index, node := range g.nodes {
	
		if node.id == id {
		
			return index, g.nodes[index]
		}
	}

	return KErrorNodeNotFound, Node{}
}

func (g *Graph) removeArcsBetweenTwoNodes(from, to uint64) {
	for index, arc := range g.arcs {
		if arc.from == from && arc.to == to {
			g.arcs = append(g.arcs[:index], g.arcs[index+1:]...)
		}
	}
}

func findNodeIdInList(list []uint64, id uint64) int {
	for idx, _ := range list {
		if list[idx] == id {
			return idx
		}
	}
	return KErrorNodeNotFound
}

func (g *Graph) findNodeIndexWithId(id uint64) int {
	for index, node := range g.nodes {
		if node.id == id {
			return index
		}
	}

	return KErrorNodeNotFound
}