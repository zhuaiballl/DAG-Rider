package DAG_Rider

import (
	"github.com/filecoin-project/go-address"
	"strings"
)

type Vertex struct {
	Round int
	Source address.Address
	Block *Block
	StrongEdges []Edge
	WeakEdges []Edge
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

