package queue

import linkedlist "github.com/maverick6912/dsa_go/linkedlist/basics/doubly-linkedlist"

// TODO: convert this queue to use generics

type DllQueue struct {
	elements  *linkedlist.DoublyLinkedList[int]
	cacheSize int
}

// New returns an initialized and empty Queue.
// requires cacheSize, if not provided defaults to zero.
func New(cacheSize int) *DllQueue {
	var dll *linkedlist.DoublyLinkedList[int]
	dll = dll.New(linkedlist.CompareDLLInt)
	return &DllQueue{elements: dll, cacheSize: cacheSize}
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (d *DllQueue) Enqueue(val int) {
	if d == nil {
		return
	}
	if d.elements.Size() > d.cacheSize {
		return
	}
	d.elements.Add(val)
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (d *DllQueue) Dequeue() int {
	if d == nil {
		return -1
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return -1
	}
	defer d.elements.Delete(0)
	val, _ := d.elements.Get(0)
	return val
}

// Peek returns the first element in the Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (d *DllQueue) Peek() int {
	if d == nil {
		return -1
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return -1
	}
	val, _ := d.elements.Get(0)
	return val
}

// Clear the queue.
func (d *DllQueue) Clear() {
	d.elements.Clear()
}

// String implements stringer interface.
func (d *DllQueue) String() string {
	return d.elements.String()
}
