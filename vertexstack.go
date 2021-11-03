package DAG_Rider

import "fmt"

type VertexStack struct {
	vertices []Vertex
}

func (st *VertexStack) Empty() bool {
	return len(st.vertices) == 0
}

func (st *VertexStack) Push(v *Vertex) error {
	st.vertices = append(st.vertices, *v)
	return nil
}

func (st *VertexStack) Pop() (v *Vertex, err error) {
	l := len(st.vertices)
	if l == 0 {
		v = nil
		err = fmt.Errorf("poping an empty VertexStack")
		return
	}
	*v = st.vertices[l-1]
	st.vertices = st.vertices[:l-1]
	err = nil
	return
}

func (st *VertexStack) Front() (*Vertex, error) {
	l := len(st.vertices)
	if l == 0 {
		return nil, fmt.Errorf("taking the front element from an empty VertexStack")
	}
	return &(st.vertices[l-1]), nil
}
