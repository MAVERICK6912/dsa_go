package queue

import (
	"fmt"
	"strings"
)

type Queue struct {
	elements []int
	size     int
}

// New returns an initialized and empty Queue.
func New() *Queue {
	return &Queue{elements: make([]int, 0)}
}

// Add values to the Queue, as per order of given values.
func (q *Queue) Add(vals ...int) {
	q.elements = append(q.elements, vals...)
	q.size += len(vals)
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (q *Queue) Enqueue(val int) {
	if q == nil {
		return
	}
	q.elements = append(q.elements, val)
	q.size += 1
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (q *Queue) Dequeue() int {
	if q == nil {
		return -1
	}
	if q.Size() == 0 || q.elements == nil {
		return -1
	}
	ret := q.elements[0]
	q.elements = q.elements[1:]
	q.size -= 1
	return ret
}

// Peek returns the first element in the Queue.
func (q *Queue) Peek() int {
	if q == nil {
		return -1
	}
	return q.elements[0]
}

// Size returns current size of Queue.
func (q *Queue) Size() int {
	return q.size
}

// Implements stringer interface for Queue
func (q *Queue) String() string {
	var sArr []string
	for _, v := range q.elements {
		sArr = append(sArr, fmt.Sprintf("%d", v))
	}
	return strings.Join(sArr, ",")
}
