package DAG_Rider

import "github.com/filecoin-project/go-address"

type Node struct {
	dag DAG
	blocksToPropose BlockQueue
	round Round
	buffer map[*Vertex]bool
	decidedWave Wave
	deliveredVertices []*Vertex
	leaderStack VertexStack
}

func (nd *Node) HasDelivered(v *Vertex) bool {
	// TODO: determine by 'delivered' mark, not searching
	for _,vv := range nd.deliveredVertices {
		if vv.Cmp(v) == 0 {
			return true
		}
	}
	return false
}

func (nd *Node) SetWeakEdges(v *Vertex, r Round) {
	v.WeakEdges = nil
	for rd := r-2; rd > 0; rd-- {
		for _, u := range nd.dag.Round[rd] {
			if !nd.dag.Path(v, u) {
				v.WeakEdges = append(v.WeakEdges, Edge{u})
			}
		}
	}
}

func (nd *Node) CreateNewVertex(r Round) (v *Vertex) {
	// wait until !nd.blocksToPropose.Empty()
	for nd.blocksToPropose.Empty() {}
	// take a block out of queue
	var err error
	v.Block, err = nd.blocksToPropose.Pop()
	if err != nil {
		panic(err)
	}
	// set round?
	//v.Round = r
	// add edges
	for _,u := range nd.dag.Round[r] {
		v.AddEdge(u)
	}
	nd.SetWeakEdges(v, r)
	return
}

func (nd *Node) AddVertex(v *Vertex) {
	nd.dag.Round[v.Round] = append(nd.dag.Round[v.Round], v)
}

func (nd *Node) VerifyEdges(v *Vertex) bool {
	for _,e := range v.StrongEdges {
		if e.To.Round <= 0 {
			return false
		}
		if !nd.dag.Exist(e.To) {
			return false
		}
	}
	return true
}

func (nd *Node) DagConstruct() {
	for {
		for v, _ := range nd.buffer {
			if v.Round <= nd.round {
				if nd.VerifyEdges(v) {
					nd.AddVertex(v)
					delete(nd.buffer, v)
				}
			}
		}
		if len(nd.dag.Round[nd.round]) >= 2*F+1 {
			if nd.round % WaveSize == 0 {
				go WaveReady(Wave(nd.round/WaveSize))
			}
			nd.round ++
			v := nd.CreateNewVertex(nd.round)
			go RBcast(v, nd.round)
		}
	}
}

func (nd *Node) GetWaveVertexLeader(w Wave) *Vertex {
	leader := ChooseLeader(w)
	r := Round((w-1) * WaveSize + 1)
	for _, v := range nd.dag.Round[r] {
		if v.Source == leader {
			return v
		}
	}
	return nil
}

func (nd *Node) OrderVertices() {
	for !nd.leaderStack.Empty() {
		v, err := nd.leaderStack.Pop()
		if err != nil { // in fact, err is always nil
			return
		}
		var verticesToDeliver []*Vertex
		for r, vertices := range nd.dag.Round {
			if r == 0 {
				continue
			}
			for _, vv := range vertices {
				if nd.dag.Path(v, vv) && !nd.HasDelivered(vv) {
					verticesToDeliver = append(verticesToDeliver, vv)
				}
			}
		}
		// TODO: make the order of elements in verticesToDeliver deterministic

		for _, vv := range verticesToDeliver {
			ADeliver(vv.Block, vv.Round, vv.Source)
			nd.deliveredVertices = append(nd.deliveredVertices, vv)
			// TODO: mark vv as delivered
		}
	}
}

func (nd *Node) WaveReady(w Wave) {
	v := nd.GetWaveVertexLeader(w)
	if v == nil {
		return
	}
	count := 0
	for _, vv := range nd.dag.Round[Round(w*WaveSize)] {
		if nd.dag.Path(vv, v){
			count ++
		}
	}
	if count <= 2*F {
		return
	}
	nd.leaderStack.Push(v)
	for ww := w-1; ww > nd.decidedWave; ww-- {
		vv := nd.GetWaveVertexLeader(ww)
		if vv != nil && nd.dag.StrongPath(v, vv) {
			nd.leaderStack.Push(vv)
			v = vv
		}
	}
	nd.decidedWave = w
	nd.OrderVertices()
}

func (nd *Node) ABcast(b *Block, r Round) {
	nd.blocksToPropose.Push(b)
}

func (nd *Node) RDeliver(v *Vertex, r Round, source address.Address) {
	v.Source = source
	v.Round = r
	if len(v.StrongEdges) >= 2*F + 1 {
		nd.buffer[v] = true
	}
}