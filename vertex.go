package DAG_Rider

import (
	"github.com/filecoin-project/go-address"
	"strings"
)

type Vertex struct {
	Round Round
	Source address.Address
	Block *Block
	StrongEdges []Edge
	WeakEdges []Edge
	delivered bool
}


func (v *Vertex) Cmp(u *Vertex) int {
	if v.Round > u.Round {
		return 1
	}
	if v.Round < u.Round {
		return -1
	}
	return strings.Compare(v.Source.String(), u.Source.String())
}

// AddEdge adds STRONG edges between vertex v and u
func (v *Vertex) AddEdge(u *Vertex) {
	v.StrongEdges = append(v.StrongEdges, Edge{u})
}

func (v *Vertex) AddWeakEdge(u *Vertex) {
	v.WeakEdges = append(v.WeakEdges, Edge{u})
}

func (v *Vertex) IsDelivered() bool {
	return v.delivered
}

func (v *Vertex) MarkDelivered() {
	v.delivered = true
}