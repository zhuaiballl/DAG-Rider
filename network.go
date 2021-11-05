package DAG_Rider

import "github.com/filecoin-project/go-address"

func RBcast(v *Vertex, r Round) {
	// TODO: finish RBcast
}

func (nd *Node) RDeliver(v *Vertex, r Round, source address.Address) {
	v.Source = source
	v.Round = r
	if len(v.StrongEdges) >= 2*F + 1 {
		nd.buffer[v] = true
	}
}

func (nd *Node) ABcast(b *Block, r Round) {
	nd.blocksToPropose.Push(b)
}

func ADeliver(b *Block, r Round, source address.Address) {
	// TODO: finish ADeliver
}
