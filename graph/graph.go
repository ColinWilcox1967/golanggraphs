package graph

const (
	KErrorNone               int = 0
	KErrorNodeNotFound       int = -1
	KErrorConflictingNodeIds int = -2
)

type Graph struct {
	nodes []Node
	arcs  []Arc
}

// package scope global variables
var (
	g Graph
)

func CreateNewGraph() Graph {

	g.nodes = nil
	g.arcs = nil

	return g
}

//
// Graph methods
//
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