package main

import (
	"fmt"
)

var (
	g Graph
)

type Arc struct {
	weight   float64
	from, to Node
}

type Node struct {
	id       uint64
	value    float64
	inbound  []Arc
	outbound []Arc
}

type Graph struct {
	nodes []Node
	arcs  []Arc
}

func (g *Graph) Init() {
	g.nodes = nil
	g.arcs = nil
}

func (g *Graph)AddNode (id uint64, value float64) int {
	if !g.uniqueId (id) {
		return -1 // define later
	} 	

	// we have a new Id
	var newNode Node 
	g.nodes = append (g.nodes, newNode)

	return 0 // KErrorNone
}

func (g *Graph)uniqueId(id uint64) bool {
	for _, node := range g.nodes {
		if node.GetId() == id {
			return false
		}
	}	
	return true
}


func (g *Graph) IsConnected(fromNode, toNode Node) bool {
	id := toNode.GetId()
	inboundArcs := fromNode.GetInboundArcList()
	outboundArcs := fromNode.GetOutboundArcList()

	for _, arc := range inboundArcs {
		if arc.GetFromNodeId() == id {
			return true
		}
	}

	for _, arc := range outboundArcs {
		if arc.GetToNodeId() == id {
			return true
		}
	}

	return false
}

func (g *Graph) GetNodeCount() int {

	return len(g.nodes)
}

func (a *Arc) GetFromNode() Node {
	return a.from
}

func (a *Arc) GetToNode() Node {
	return a.to
}

func (a *Arc) GetFromNodeId() uint64 {
	n := a.GetFromNode()
	return n.GetId()
}

func (a *Arc) GetToNodeId() uint64 {
	n := a.GetToNode()
	return n.GetId()
}

func (n *Node) GetInboundArcCount() int {
	return len(n.inbound)
}

func (n *Node) GetOutboundArcCount() int {
	return len(n.outbound)
}

func (n *Node) GetId() uint64 {
	return n.id
}

func (n *Node) GetInboundArcList() []Arc {
	return n.inbound
}

func (n *Node) GetOutboundArcList() []Arc {
	return n.outbound
}

func main() {

	g.Init()
	fmt.Printf("Number of nodes = %d\n", g.GetNodeCount())
}
