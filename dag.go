package DAG_Rider

type Edge struct {
	To *Vertex
}

type DAG struct {
	Vertices []Vertex
}

// Path checks if it is possible to go from v to u through strong or weak edges
func (dag *DAG) Path(v,u *Vertex) bool {
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