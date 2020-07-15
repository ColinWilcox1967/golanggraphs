package graph

const (
	KErrorNone               int = 0
	KErrorNodeNotFound       int = -1
	KErrorConflictingNodeIds int = -2
)

type Arc struct {
	weight   float64
	from, to uint64
}

type Node struct {
	id       uint64
	value    float64
	inbound  []uint64
	outbound []uint64
}

type Graph struct {
	nodes []Node
	arcs  []Arc
}

var (
	g Graph
)

func CreateNewGraph() Graph {

	g.nodes = nil
	g.arcs = nil

	return g
}

//
// Arc methods
//
func (a *Arc) New(fromNode, toNode uint64, weight float64) Arc {

	var newArc Arc
	newArc.from = fromNode
	newArc.to = toNode
	newArc.weight = weight

	return newArc
}

func (a *Arc) GetFromNode() Node {
	status, node := g.findNodeWithId(a.from)
	if status == 0 {
		return node
	}
	return Node{}
}

func (a *Arc) GetToNode() Node {
	status, node := g.findNodeWithId(a.to)
	if status == 0 {
		return node
	}

	return Node{}
}

func (a *Arc) GetFromNodeId() uint64 {
	node := a.GetFromNode()
	return node.id
}

func (a *Arc) GetToNodeId() uint64 {
	node := a.GetToNode()
	return node.id
}

func (a *Arc) GetWeight() float64 {
	return a.weight
}

//
// Node methods
//
func (n *Node) New(id uint64, value float64) Node {
	var node Node
	node.value = value
	node.id = id
	node.inbound = nil
	node.outbound = nil
	return node
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

func (n *Node) GetValue() float64 {
	return n.value
}

func (n *Node) GetInboundArcList() []uint64 {
	return n.inbound
}

func (n *Node) GetOutboundArcList() []uint64 {
	return n.outbound
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

func (g *Graph) uniqueId(id uint64) bool {
	for _, node := range g.nodes {
		if node.GetId() == id {
			return false
		}
	}
	return true
}

func (g *Graph) GetNodeInboundConnections(index int) []uint64 {
	return g.nodes[index].inbound
}

func (g *Graph) GetNodeOutboundConnections(index int) []uint64 {
	return g.nodes[index].outbound
}

func (g *Graph) IsConnected(fromNode, toNode uint64) bool {

	status, node := g.findNodeWithId(toNode)
	if status == -1 {
		return false
	}

	inboundArcs := node.GetInboundArcList()
	outboundArcs := node.GetOutboundArcList()

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

func (g *Graph) GetNodeCount() int {

	return len(g.nodes)
}

func (g *Graph) GetArcCount() int {
	return len(g.arcs)
}

// private helper function
func (g *Graph) findNodeWithId(id uint64) (int, Node) {
	for index, node := range g.nodes {
		if node.id == id {
			return index, g.nodes[index]
		}
	}

	return KErrorNodeNotFound, Node{}
}

func (g *Graph) findNodeIndexWithId(id uint64) int {
	for index, node := range g.nodes {
		if node.id == id {
			return index
		}
	}

	return KErrorNodeNotFound
}
