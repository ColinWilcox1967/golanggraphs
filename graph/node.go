package graph

type Node struct {
	id       uint64
	value    float64
	inbound  []uint64
	outbound []uint64
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

func (n *Node) InboundArcCount() int {
	return len(n.inbound)
}

func (n *Node) OutboundArcCount() int {
	return len(n.outbound)
}

func (n *Node) Id() uint64 {
	return n.id
}

func (n *Node) Value() float64 {
	return n.value
}

func (n *Node) InboundArcList() []uint64 {
	return n.inbound
}

func (n *Node) OutboundArcList() []uint64 {
	return n.outbound
}


	
	