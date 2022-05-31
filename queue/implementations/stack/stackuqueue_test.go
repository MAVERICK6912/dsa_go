package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackUQueuePush(t *testing.T) {
	s := New()

	s.Push(20)
	assert.Equal(t, 20, s.Peek())
	assert.Equal(t, 1, s.Size())

	s.Push(69)
	assert.Equal(t, 69, s.Peek())
	assert.Equal(t, 2, s.Size())

	s.Push(9)
	assert.Equal(t, 9, s.Peek())
	assert.Equal(t, 3, s.Size())
}

func TestStackUQueuePop(t *testing.T) {
	s := New()

	s.Push(125)
	s.Push(2169)
	s.Push(3)

	val := s.Pop()
	assert.Equal(t, 3, val)
	assert.Equal(t, 2, s.Size())

	val = s.Pop()
	assert.Equal(t, 2169, val)
	assert.Equal(t, 1, s.Size())

	val = s.Pop()
	assert.Equal(t, 125, val)
	assert.Equal(t, 0, s.Size())

	// popping when stack is empty
	val = s.Pop()
	assert.Equal(t, -1, val)
	assert.Equal(t, 0, s.Size())
}
