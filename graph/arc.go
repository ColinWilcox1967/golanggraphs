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

func (a *Arc) FromNode() Node {
	status, node := g.findNodeWithId(a.from)
	if status == 0 {
		return node
	}
	return Node{}
}

func (a *Arc) ToNode() Node {
	status, node := g.findNodeWithId(a.to)
	if status == 0 {
		return node
	}

	return Node{}
}

func (a *Arc) FromNodeId() uint64 {
	node := a.FromNode()
	return node.id
}

func (a *Arc) ToNodeId() uint64 {
	node := a.ToNode()
	return node.id
}

func (a *Arc) Weight() float64 {
	return a.weight
}


