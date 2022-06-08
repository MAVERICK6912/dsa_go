package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLevelOrder(t *testing.T) {
	bt := New()

	vals := []int{21, 3, 54, 69, 2, 1, 7}

	bt.AddLevelOrder(vals...)

	assert.Equal(t, bt.root.data, 21)
	assert.Equal(t, bt.root.left.data, 3)
	assert.Equal(t, bt.root.right.data, 54)
	assert.Equal(t, bt.root.left.left.data, 69)
	assert.Equal(t, bt.root.left.right.data, 2)
	assert.Equal(t, bt.root.right.left.data, 1)
	assert.Equal(t, bt.root.right.right.data, 7)

}
