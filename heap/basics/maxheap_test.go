package heap

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var m *MaxHeap[int]
	m = m.New(utils.IntComparator)
	vals := []int{25, 3, 67, 4, 5, 13}

	for _, v := range vals {
		m.Insert(v)
	}
	/*
				67
			   /  \
			 25    13
			/ \   /
		   5   4 3
	*/
	assert.Equal(t, 67, m.elems[0])
	assert.Equal(t, len(vals), m.Length())
}

func TestExtract(t *testing.T) {
	var m *MaxHeap[int]
	m = m.New(utils.IntComparator)
	vals := []int{25, 3, 67, 4, 5, 13}

	for _, v := range vals {
		m.Insert(v)
	}

	/*
				67
			   /  \
			 25    13
			/ \   /
		   5   4 3
	*/

	actual, err := m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 67, actual)
	assert.Equal(t, len(vals)-1, m.Length())

	/*
				25
			   /  \
			 5    13
			/ \
		   3   4
	*/

	actual, err = m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 25, actual)
	assert.Equal(t, len(vals)-2, m.Length())
	/*
				13
			   /  \
			 5     4
			/
		   3
	*/

	actual, err = m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 13, actual)
	assert.Equal(t, len(vals)-3, m.Length())
	/*
			5
		   /  \
		 3     4
	*/

	actual, err = m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 5, actual)
	assert.Equal(t, len(vals)-4, m.Length())
	/*
			4
		   / \
		nil   3
	*/

	actual, err = m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 4, actual)
	assert.Equal(t, len(vals)-5, m.Length())
	/*
			3
		   / \
		nil   nil
	*/
	actual, err = m.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 3, actual)
	assert.Equal(t, len(vals)-6, m.Length())

	_, err = m.Extract()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NoElements)
	}
}
