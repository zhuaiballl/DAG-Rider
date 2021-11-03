package DAG_Rider

type Round int
type Wave int

type DAG struct {
	Round map[Round][]*Vertex
}

func (dag *DAG) Exist(v *Vertex) bool {
	for _,ver := range dag.Round[v.Round] {
		if v.Cmp(ver) == 0 {
			return true
		}
	}
	return false
}

// Path checks if it is possible to go from v to u through strong or weak edges
func (dag *DAG) Path(v,u *Vertex) bool {
	if !dag.Exist(v) || !dag.Exist(u) {
		return false
	}
	// TODO: change DFS to BFS or add memory to it
	if v.Round < u.Round {
		return false
	}
	if v.Round == u.Round {
		if v.Cmp(u) == 0 {
			return true
		}else{
			return false
		}
	}
	for _,se := range v.StrongEdges {
		if dag.Path(se.To, u) {
			return true
		}
	}
	for _,we := range v.WeakEdges {
		if dag.Path(we.To, u) {
			return true
		}
	}
	return false
}

// StrongPath checks if it is possible to go from v to u through strong edges
func (dag *DAG) StrongPath(v,u *Vertex) bool {
	if !dag.Exist(v) || !dag.Exist(u) {
		return false
	}
	// TODO: change DFS to BFS or add memory to it
	if v.Round < u.Round {
		return false
	}
	if v.Round == u.Round {
		if v.Cmp(u) == 0 {
			return true
		}else{
			return false
		}
	}
	for _,se := range v.StrongEdges {
		if dag.StrongPath(se.To, u) {
			return true
		}
	}
	return false
}