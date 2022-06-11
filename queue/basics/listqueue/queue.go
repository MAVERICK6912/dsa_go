package queue

import (
	"fmt"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
)

type Queue[T any] struct {
	cmp      func(T, T) int
	elements []T
	size     int
}

// New returns an initialized and empty Queue.
func (q *Queue[T]) New(cmp func(T, T) int) *Queue[T] {
	return &Queue[T]{elements: make([]T, 0), cmp: cmp}
}

// Add values to the Queue, as per order of given values.
func (q *Queue[T]) Add(vals ...T) {
	q.elements = append(q.elements, vals...)
	q.size += len(vals)
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (q *Queue[T]) Enqueue(val T) {
	if q == nil {
		return
	}
	q.elements = append(q.elements, val)
	q.size += 1
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (q *Queue[T]) Dequeue() (T, error) {
	var ret T
	if q == nil {
		return ret, errors.UninitializedError
	}
	if q.Size() == 0 || q.elements == nil {
		return ret, errors.UninitializedError
	}
	ret = q.elements[0]
	q.elements = q.elements[1:]
	q.size -= 1
	return ret, nil
}

// Peek returns the first element in the Queue.
func (q *Queue[T]) Peek() (T, error) {
	if q == nil {
		var ret T
		return ret, errors.UninitializedError
	}
	return q.elements[0], nil
}

// Size returns current size of Queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// Implements stringer interface for Queue
func (q *Queue[T]) String() string {
	var sArr []string
	for _, v := range q.elements {
		sArr = append(sArr, fmt.Sprintf("%v", v))
	}
	return strings.Join(sArr, ",")
}
