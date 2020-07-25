package graph

type Arc struct {
	from, to uint64
	weight   float64
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

func (a *Arc) FromNode(g *Graph) Node {

	status, node := g.findNodeWithId(a.from)
	if status >= 0 {
		return node
	}
	return Node{}
}

func (a *Arc) ToNode(g *Graph) Node {

	status, node := g.findNodeWithId(a.to)

	if status >= 0 {
		return node
	}

	return Node{}
}

func (a *Arc) FromNodeId(g *Graph) uint64 {
	node := a.FromNode(g)
	return node.id
}

func (a *Arc) ToNodeId(g *Graph) uint64 {
	node := a.ToNode(g)
	return node.id
}

func (a *Arc) SetFromNodeId(fromNode uint64) {
	a.from = fromNode
}

func (a *Arc) SetToNodeId(toNode uint64) {
	a.to = toNode
}

func (a *Arc) Weight() float64 {
	return a.weight
}

func (a *Arc) SetWeight(weight float64) {
	a.weight = weight
}
