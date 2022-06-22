package graph

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

var (
	verticesInt = []int{98, 5, 36, 1, 69, 7}
	edgesInt    = []struct {
		from int
		to   int
	}{
		{
			from: 98,
			to:   1,
		},
		{
			from: 5,
			to:   1,
		},
		{
			from: 36,
			to:   5,
		},
		{
			from: 69,
			to:   98,
		},
		{
			from: 7,
			to:   36,
		},
		{
			from: 5,
			to:   7,
		},
	}
)

func TestAddVertex(t *testing.T) {
	var g *DirectedGraph[int]
	g = g.New(utils.IntComparator)

	for _, v := range verticesInt {
		err := g.AddVertex(v)
		if err != nil {
			t.Error("error while adding vertex to graph, error was: ", err.Error())
		}
	}
	actual, err := g.Values()
	if err != nil {
		t.Error("error while getting vertices from graph, error was: ", err.Error())
	}
	assert.ElementsMatch(t, verticesInt, actual)

	// inserting a key that already exist in graph
	err = g.AddVertex(1)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.KeyAlreadyExist)
	}
}

func TestAddEdge(t *testing.T) {
	var g *DirectedGraph[int]
	g = g.New(utils.IntComparator)

	for _, v := range verticesInt {
		err := g.AddVertex(v)
		if err != nil {
			t.Error("error while adding vertex to graph, error was: ", err.Error())
		}
	}
	for _, e := range edgesInt {
		err := g.AddEdge(e.from, e.to)
		if err != nil {
			t.Error("error while adding vertex to graph, error was: ", err.Error())
		}
	}
	for _, e := range edgesInt {
		ok := g.IsConnected(e.from, e.to)
		if !ok {
			t.Errorf("vertex %d and %d are not connected ", e.from, e.to)
		}
	}
	// making a connection to a vertex that doesn't exist
	err := g.AddEdge(0, 9)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
	// making a connection to a vertex that already exist
	err = g.AddEdge(7, 36)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.KeyAlreadyExist)
	}
}

func TestRemoveVertex(t *testing.T) {
	var g *DirectedGraph[int]
	g = g.New(utils.IntComparator)

	for _, v := range verticesInt {
		err := g.AddVertex(v)
		if err != nil {
			t.Error("error while adding vertex to graph, error was: ", err.Error())
		}
	}
	for _, e := range edgesInt {
		err := g.AddEdge(e.from, e.to)
		if err != nil {
			t.Error("error while adding vertex to graph, error was: ", err.Error())
		}
	}

	err := g.RemoveVertex(verticesInt[3])
	if err != nil {
		t.Errorf("error while removing vertex with key %d", verticesInt[3])
	}
	actual := g.getVertex(verticesInt[3])
	assert.Nil(t, actual)

	// removing a vertex that doesn't exist
	err = g.RemoveVertex(99)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}
