package queue

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maverick6912/dsa_go/utils"
)

type PriorityQueue struct {
	Queue []Element
	size  int
}

type Element struct {
	Priority int
	value    int
}

// New returns an initialized and empty PriorityQueue.
func New() *PriorityQueue {
	return &PriorityQueue{Queue: make([]Element, 0)}
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (p *PriorityQueue) Enqueue(val, pri *int) {
	if pri == nil {
		return
	}
	if p == nil {
		return
	}
	if p.Queue == nil {
		return
	}
	p.Queue = append(p.Queue, Element{
		value:    *val,
		Priority: *pri,
	})
	p.size += 1
	// utils.Sort(p.Queue, utils.PriorityComparator)
	sort.Slice(p.Queue, p.Less)
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (p *PriorityQueue) Dequeue() int {
	if p == nil {
		return -1
	}
	if p.Queue == nil || p.Size() == 0 {
		return -1
	}
	val := p.Queue[0].value
	p.Queue = p.Queue[1:]
	p.size -= 1
	return val
}

// Peek returns the first element in the Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (p *PriorityQueue) Peek() int {
	if p == nil {
		return -1
	}
	if p.Queue == nil {
		return -1
	}
	return p.Queue[0].value
}

// Size returns current size of queue.
func (p *PriorityQueue) Size() int {
	return p.size
}

// String implements stringer interface.
func (p *PriorityQueue) String() string {
	sArr := make([]string, 0)
	for _, v := range p.Queue {
		sArr = append(sArr, fmt.Sprintf("%d", v.value))
	}
	return strings.Join(sArr, ",")
}

func (p *PriorityQueue) Less(i, j int) bool {

	return utils.IntComparator(p.Queue[i].Priority, p.Queue[j].Priority) < 0

}
