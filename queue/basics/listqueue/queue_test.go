package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEnqueue(t *testing.T) {
	q := New()

	q.Enqueue(34)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 34, q.Peek())

	q.Enqueue(45)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 34, q.Peek())
}

func TestQueueDequeue(t *testing.T) {
	vals := []int{7, 8, 9}

	q := New()
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)

	val := q.Dequeue()
	assert.Equal(t, len(vals)-1, q.Size())
	assert.Equal(t, vals[0], val)

	val = q.Dequeue()
	assert.Equal(t, len(vals)-2, q.Size())
	assert.Equal(t, vals[1], val)

	val = q.Dequeue()
	assert.Equal(t, len(vals)-3, q.Size())
	assert.Equal(t, vals[2], val)
	// Dequeuing when queue is empty
	val = q.Dequeue()
	assert.Equal(t, len(vals)-3, q.Size())
	assert.Equal(t, -1, val)
}
