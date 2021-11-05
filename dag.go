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
	vis := make(map[*Vertex]bool)
	st := VertexStack{}
	st.Push(v)
	vis[v]=true
	for !st.Empty() {
		cur,_ := st.Pop()
		for _,se := range cur.StrongEdges {
			to := se.To
			if _,fd := vis[to]; !fd {
				if u.Cmp(to) == 0 {
					return true
				}
				if to.Round > u.Round {
					st.Push(to)
				}
			}
		}
		for _,we := range cur.StrongEdges {
			to := we.To
			if _,fd := vis[to]; !fd {
				if u.Cmp(to) == 0 {
					return true
				}
				if to.Round > u.Round {
					st.Push(to)
				}
			}
		}
	}
	return false
}

// StrongPath checks if it is possible to go from v to u through strong edges
func (dag *DAG) StrongPath(v,u *Vertex) bool {
	if !dag.Exist(v) || !dag.Exist(u) {
		return false
	}
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
	vis := make(map[*Vertex]bool)
	st := VertexStack{}
	st.Push(v)
	vis[v]=true
	for !st.Empty() {
		cur,_ := st.Pop()
		for _,se := range cur.StrongEdges {
			to := se.To
			if _,fd := vis[to]; !fd {
				if u.Cmp(to) == 0 {
					return true
				}
				if to.Round > u.Round {
					st.Push(to)
				}
			}
		}
	}
	for _,se := range v.StrongEdges {
		if dag.StrongPath(se.To, u) {
			return true
		}
	}
	return false
}