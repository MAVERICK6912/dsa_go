package queue

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	linkedlist "github.com/maverick6912/dsa_go/linkedlist/basics/doubly-linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedListQueueEnqueue(t *testing.T) {
	t.Parallel()
	var dllQ DllQueue[int]
	dllQ = *dllQ.New(3, linkedlist.CompareDLLInt)

	err := dllQ.Enqueue(59)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}
	actual, err := dllQ.Peek()
	if err != nil {
		t.Error("error while peeking from queue")
	}
	assert.Equal(t, 59, actual)
	assert.Equal(t, 1, dllQ.elements.Size())

	err = dllQ.Enqueue(69)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}
	assert.Equal(t, 2, dllQ.elements.Size())

	err = dllQ.Enqueue(369)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}
	assert.Equal(t, 3, dllQ.elements.Size())
}

func TestDoublyLinkedListQueueDequeue(t *testing.T) {
	t.Parallel()
	var dllQ DllQueue[int]
	dllQ = *dllQ.New(3, linkedlist.CompareDLLInt)

	err := dllQ.Enqueue(59)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}
	err = dllQ.Enqueue(69)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}
	err = dllQ.Enqueue(369)
	if err != nil {
		t.Error("error while enqueuing to queue")
	}

	actual, err := dllQ.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from queue")
	}
	assert.Equal(t, 59, actual)
	assert.Equal(t, 2, dllQ.elements.Size())

	actual, err = dllQ.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from queue")
	}
	assert.Equal(t, 69, actual)
	assert.Equal(t, 1, dllQ.elements.Size())

	actual, err = dllQ.Dequeue()
	if err != nil {
		t.Error("error while dequeuing from queue")
	}
	assert.Equal(t, 369, actual)
	assert.Equal(t, 0, dllQ.elements.Size())
	// Dequeuing when queue is empty
	_, err = dllQ.Dequeue()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
	assert.Equal(t, 0, dllQ.elements.Size())
}
