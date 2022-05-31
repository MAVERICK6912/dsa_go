package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListQueueEnqueue(t *testing.T) {
	dllQ := New(3)

	dllQ.Enqueue(59)
	assert.Equal(t, 59, dllQ.Peek())
	assert.Equal(t, 1, dllQ.elements.Size())

	dllQ.Enqueue(69)
	assert.Equal(t, 2, dllQ.elements.Size())

	dllQ.Enqueue(369)
	assert.Equal(t, 3, dllQ.elements.Size())
}

func TestDoublyLinkedListQueueDequeue(t *testing.T) {
	dllQ := New(3)
	dllQ.Enqueue(59)
	dllQ.Enqueue(69)
	dllQ.Enqueue(369)

	res := dllQ.Dequeue()
	assert.Equal(t, 59, res)
	assert.Equal(t, 2, dllQ.elements.Size())

	res = dllQ.Dequeue()
	assert.Equal(t, 69, res)
	assert.Equal(t, 1, dllQ.elements.Size())

	res = dllQ.Dequeue()
	assert.Equal(t, 369, res)
	assert.Equal(t, 0, dllQ.elements.Size())
	// Dequeuing when queue is empty
	res = dllQ.Dequeue()
	assert.Equal(t, -1, res)
	assert.Equal(t, 0, dllQ.elements.Size())
}
