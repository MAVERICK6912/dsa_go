package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPriorityQueueEnqueue(t *testing.T) {
	pq := New()

	val, pri := 23, 5
	pq.Enqueue(&val, &pri)
	assert.Equal(t, val, pq.Peek())
	assert.Equal(t, 1, pq.Size())

	val, pri = 2, 1
	pq.Enqueue(&val, &pri)
	assert.Equal(t, 2, pq.Size())
	assert.Equal(t, val, pq.Peek())

	val, pri = 1, 0
	pq.Enqueue(&val, &pri)
	assert.Equal(t, 3, pq.Size())
	assert.Equal(t, val, pq.Peek())
	// Enqueuing element with same priority
	val, pri = 36, 5
	pq.Enqueue(&val, &pri)
	assert.Equal(t, 4, pq.Size())
}

func TestListPriorityQueueDequeue(t *testing.T) {
	pq := New()

	val, pri := 23, 5
	pq.Enqueue(&val, &pri)
	val, pri = 0, 0
	pq.Enqueue(&val, &pri)
	val, pri = 3436, 1
	pq.Enqueue(&val, &pri)
	val, pri = 6969, 5
	pq.Enqueue(&val, &pri)

	val = pq.Dequeue()
	assert.Equal(t, 0, val)

	val = pq.Dequeue()
	assert.Equal(t, 3436, val)

	val = pq.Dequeue()
	assert.Equal(t, 23, val)

	val = pq.Dequeue()
	assert.Equal(t, 6969, val)
	// Dequeuing when queue is empty
	val = pq.Dequeue()
	assert.Equal(t, -1, val)
}
