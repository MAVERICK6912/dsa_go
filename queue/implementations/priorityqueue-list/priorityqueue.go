package queue

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
)

type PriorityQueue[T any] struct {
	cmp   func(T, T) int
	Queue []Element[T]
	size  int
}

type Element[T any] struct {
	Priority int
	value    T
}

// New returns an initialized and empty PriorityQueue.
func (p *PriorityQueue[T]) New(cmp func(T, T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{Queue: make([]Element[T], 0), cmp: cmp}
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (p *PriorityQueue[T]) Enqueue(val T, pri *int) error {
	if pri == nil {
		return errors.MissingPriority
	}
	if p == nil {
		return errors.UninitializedError
	}
	if p.Queue == nil {
		return errors.UninitializedError
	}
	p.Queue = append(p.Queue, Element[T]{
		value:    val,
		Priority: *pri,
	})
	p.size += 1
	// utils.Sort(p.Queue, utils.PriorityComparator)
	sort.Slice(p.Queue, p.Less)
	return nil
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (p *PriorityQueue[T]) Dequeue() (T, error) {
	var ret T
	if p == nil {
		return ret, errors.UninitializedError
	}
	if p.Queue == nil || p.Size() == 0 {
		return ret, errors.UninitializedError
	}
	ret = p.Queue[0].value
	p.Queue = p.Queue[1:]
	p.size -= 1
	return ret, nil
}

// Peek returns the first element in the Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (p *PriorityQueue[T]) Peek() (T, error) {
	var ret T
	if p == nil {
		return ret, errors.UninitializedError
	}
	if p.Queue == nil {
		return ret, errors.UninitializedError
	}
	ret = p.Queue[0].value
	return ret, nil
}

// Size returns current size of queue.
func (p *PriorityQueue[T]) Size() int {
	return p.size
}

// String implements stringer interface.
func (p *PriorityQueue[T]) String() string {
	sArr := make([]string, 0)
	for _, v := range p.Queue {
		sArr = append(sArr, fmt.Sprintf("%v", v.value))
	}
	return strings.Join(sArr, ",")
}

func (p *PriorityQueue[T]) Less(i, j int) bool {
	return utils.IntComparator(p.Queue[i].Priority, p.Queue[j].Priority) < 0
}
