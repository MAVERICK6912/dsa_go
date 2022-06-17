package heap

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var h *Heap[int]
	h = h.New(utils.IntComparator)
	vals := []int{25, 3, 67, 4, 5, 13}

	for _, v := range vals {
		h.Insert(v)
	}
	/*
				67
			   /  \
			 25    13
			/ \   /
		   5   4 3
	*/
	assert.Equal(t, 67, h.elems[0])
	assert.Equal(t, len(vals), h.Length())
}

func TestExtract(t *testing.T) {
	var h *Heap[int]
	h = h.New(utils.IntComparator)
	vals := []int{25, 3, 67, 4, 5, 13}

	for _, v := range vals {
		h.Insert(v)
	}

	/*
				67
			   /  \
			 25    13
			/ \   /
		   5   4 3
	*/

	actual, err := h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 67, actual)
	assert.Equal(t, len(vals)-1, h.Length())

	/*
				25
			   /  \
			 5    13
			/ \
		   3   4
	*/

	actual, err = h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 25, actual)
	assert.Equal(t, len(vals)-2, h.Length())
	/*
				13
			   /  \
			 5     4
			/
		   3
	*/

	actual, err = h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 13, actual)
	assert.Equal(t, len(vals)-3, h.Length())
	/*
			5
		   /  \
		 3     4
	*/

	actual, err = h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 5, actual)
	assert.Equal(t, len(vals)-4, h.Length())
	/*
			4
		   / \
		nil   3
	*/

	actual, err = h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 4, actual)
	assert.Equal(t, len(vals)-5, h.Length())
	/*
			3
		   / \
		nil   nil
	*/
	actual, err = h.Extract()
	assert.Nil(t, err)
	assert.Equal(t, 3, actual)
	assert.Equal(t, len(vals)-6, h.Length())

	_, err = h.Extract()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NoElements)
	}
}
