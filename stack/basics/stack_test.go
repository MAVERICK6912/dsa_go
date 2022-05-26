package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	vals := []int{24, 69, 9, 12, 2, 56}
	var s Stack

	s.Add(vals...)

	s.Push(35)
	assert.Equal(t, 35, s.Peek())
	assert.Equal(t, len(vals)+1, s.Size())

	s.Pop()
	assert.Equal(t, 56, s.Peek())
	assert.Equal(t, len(vals), s.Size())
}
