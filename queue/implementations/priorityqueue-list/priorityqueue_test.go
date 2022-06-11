package queue

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestListPriorityQueueEnqueue(t *testing.T) {
	t.Parallel()
	var pq *PriorityQueue[int]
	pq = pq.New(utils.IntComparator)

	val, pri := 23, 5
	err := pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	actual, err := pq.Peek()
	if err != nil {
		t.Error("error while peeking from priority queue")
	}
	assert.Equal(t, val, actual)
	assert.Equal(t, 1, pq.Size())

	val, pri = 2, 1
	err = pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	assert.Equal(t, 2, pq.Size())
	actual, err = pq.Peek()
	if err != nil {
		t.Error("error while peeking from priority queue")
	}
	assert.Equal(t, val, actual)

	val, pri = 1, 0
	err = pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	assert.Equal(t, 3, pq.Size())
	actual, err = pq.Peek()
	if err != nil {
		t.Error("error while peeking from priority queue")
	}
	assert.Equal(t, val, actual)
	// Enqueuing element with same priority
	val, pri = 36, 5
	err = pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	assert.Equal(t, 4, pq.Size())
}

func TestListPriorityQueueDequeue(t *testing.T) {
	t.Parallel()
	var pq *PriorityQueue[int]
	pq = pq.New(utils.IntComparator)

	val, pri := 23, 5
	pq.Enqueue(val, &pri)

	val, pri = 0, 0
	err := pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	val, pri = 3436, 1
	err = pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}
	val, pri = 6969, 5
	err = pq.Enqueue(val, &pri)
	if err != nil {
		t.Error("error while enqueuing to priority queue")
	}

	actual, err := pq.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from priority queue")
	}
	assert.Equal(t, 0, actual)

	actual, err = pq.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from priority queue")
	}
	assert.Equal(t, 3436, actual)

	actual, err = pq.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from priority queue")
	}
	assert.Equal(t, 23, actual)

	actual, err = pq.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from priority queue")
	}
	assert.Equal(t, 6969, actual)
	// Dequeuing when queue is empty
	_, err = pq.Dequeue()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
}
