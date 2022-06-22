package graph

import (
	"github.com/maverick6912/dsa_go/errors"
)

type DirectedGraph[T any] struct {
	vertices []*DirectedVertex[T]
	cmp      func(T, T) int
}

type DirectedVertex[T any] struct {
	key      T
	adjacent []*DirectedVertex[T]
}

// New initiaizes the passed graph and returns it.
func (g *DirectedGraph[T]) New(cmp func(T, T) int) *DirectedGraph[T] {
	g = &DirectedGraph[T]{cmp: cmp}
	return g
}

// New initializes the passed vertex and returns it.
func (v *DirectedVertex[T]) New() *DirectedVertex[T] {
	v = &DirectedVertex[T]{}
	return v
}

// AddVertex insert a vertex to graph with the given key.
// returns uninitialized error if graph is not initialized.
func (g *DirectedGraph[T]) AddVertex(k T) error {
	if g == nil {
		return errors.UninitializedError
	}
	if g.contains(g.vertices, k) {
		return errors.KeyAlreadyExist
	}
	g.vertices = append(g.vertices, &DirectedVertex[T]{key: k})
	return nil
}

// contains checks if a key exists in the given list of vertices.
// returns true if key is found and false if not found.
func (g DirectedGraph[T]) contains(sv []*DirectedVertex[T], k T) bool {
	for _, v := range sv {
		if g.cmp(v.key, k) == 0 {
			return true
		}
	}
	return false
}

// AddEdge adds an edge given the to and from vertex key.
// returns notFound error if the from or to vertices don't exist
// and KeyAlreadyExist error if the edge already exist between from and to vertices.
func (g *DirectedGraph[T]) AddEdge(fromK, toK T) error {
	// get vertex
	fromVertex, toVertex := g.getVertex(fromK), g.getVertex(toK)
	if fromVertex == nil || toVertex == nil {
		return errors.NotFound
	}
	// check if edge already exist
	if g.contains(fromVertex.adjacent, toK) {
		return errors.KeyAlreadyExist
	}
	// add edge
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	return nil
}

// getVertex gets the vertex for the specified key.
func (g DirectedGraph[T]) getVertex(k T) *DirectedVertex[T] {
	for _, v := range g.vertices {
		if g.cmp(v.key, k) == 0 {
			return v
		}
	}
	return nil
}

// IsConnected returns true if passed from and to key are connected, false otherwise.
func (g DirectedGraph[T]) IsConnected(fromK, toK T) bool {
	v := g.getVertex(fromK)
	if v == nil {
		return false
	}
	return g.contains(v.adjacent, toK)
}

func (g DirectedGraph[T]) Values() ([]T, error) {
	ret := make([]T, 0)
	if g.vertices == nil {
		return ret, errors.UninitializedError
	}
	for _, v := range g.vertices {
		ret = append(ret, v.key)
	}
	return ret, nil
}

/*
DirectedGraph g is an adjecency list:
g: [{v1,adjv1:[v2,v4] },{v2,adjv2: [v1]},{v3,adjv3: [v2]},{v4,adjv4:[v2,v3]}]

if we want to remove a vertex, then:
 - we need to first remove all associations from the connecting vertices:
	- check in all vertices adjacency list.
	- if vertex to remove exist, then remove it from adjacency list of that vertex.
 - remove the vertex from graph
*/
// RemoveVertex removes a vertex from di-graph if it exist in the graph.
// returns uninitialzed error if graph is not initialized
// and returns NotFound error if vertex doesn't exist in the graph.
func (g *DirectedGraph[T]) RemoveVertex(k T) error {
	if g == nil {
		return errors.UninitializedError
	}
	if !g.contains(g.vertices, k) {
		return errors.NotFound
	}
	// removing associations from adjoining vertices.
	for _, v := range g.vertices {
		if g.contains(v.adjacent, k) {
			v.removeFromList(k, g.cmp)
		}
	}
	// removing vertex from graph.
	g.removeFromList(k)
	return nil
}

// removeFromList removes an element in given vertex adjacencyList.
func (v *DirectedVertex[T]) removeFromList(kR T, cmp func(T, T) int) {
	var t []*DirectedVertex[T]
	for _, ver := range v.adjacent {
		if cmp(ver.key, kR) != 0 {
			t = append(t, ver)
		}
	}
	v.adjacent = t
}

// removeFromList removes an element in given vertex adjacencyList.
func (g *DirectedGraph[T]) removeFromList(kR T) {
	var t []*DirectedVertex[T]
	for _, ver := range g.vertices {
		if g.cmp(ver.key, kR) != 0 {
			t = append(t, ver)
		}
	}
	g.vertices = t
}

// Vertices returns all vretices in a di-graph.
// returns UninitializedError if di-graph is uninitialized.
func (g DirectedGraph[T]) Vertices() ([]*DirectedVertex[T], error) {
	if g.vertices == nil {
		return nil, errors.UninitializedError
	}
	return g.vertices, nil
}