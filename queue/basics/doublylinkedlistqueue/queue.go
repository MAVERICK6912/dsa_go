package queue

import linkedlist "github.com/maverick6912/dsa_go/linkedlist/basics/doubly-linkedlist"

// TODO: add queue based on dll.

type DllQueue struct {
	elements  *linkedlist.DoublyLinkedList
	cacheSize int
}

func New(cacheSize int) *DllQueue {
	return &DllQueue{elements: linkedlist.New(), cacheSize: cacheSize}
}

func (d *DllQueue) Enqueue(val int) {
	if d == nil {
		return
	}
	if d.elements.Size() > d.cacheSize {
		return
	}
	d.elements.Add(val)
}

func (d *DllQueue) Dequeue() int {
	if d == nil {
		return -1
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return -1
	}
	defer d.elements.Delete(0)
	return d.elements.Get(0)
}

func (d *DllQueue) Peek() int {
	if d == nil {
		return -1
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return -1
	}
	return d.elements.Get(0)
}

func (d *DllQueue) Clear() {
	d.elements.Clear()
}

func (d *DllQueue) String() string {
	return d.elements.String()
}
