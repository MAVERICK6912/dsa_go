package queue

import (
	"github.com/maverick6912/dsa_go/errors"
	linkedlist "github.com/maverick6912/dsa_go/linkedlist/basics/doubly-linkedlist"
)

// TODO: convert this queue to use generics

type DllQueue[T any] struct {
	elements  *linkedlist.DoublyLinkedList[T]
	cacheSize int
}

// New returns an initialized and empty Queue.
// requires cacheSize, if not provided defaults to zero.
func (d *DllQueue[T]) New(cacheSize int, cmp func(*linkedlist.DLLNode[T], *linkedlist.DLLNode[T]) int) *DllQueue[T] {
	var dll *linkedlist.DoublyLinkedList[T]
	dll = dll.New(cmp)
	return &DllQueue[T]{elements: dll, cacheSize: cacheSize}
}

// Enqueue adds given value to end of Queue.
// Does nothing if `Queue` is not initialized
func (d *DllQueue[T]) Enqueue(val T) error {
	if d == nil {
		return errors.UninitializedError
	}
	if d.elements.Size() > d.cacheSize {
		return errors.IndexOutOfBound
	}
	d.elements.Add(val)
	return nil
}

// Dequeue deletes and returns first element from Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (d *DllQueue[T]) Dequeue() (T, error) {
	var ret T
	if d == nil {
		return ret, errors.UninitializedError
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return ret, errors.UninitializedError
	}
	defer d.elements.Delete(0)
	ret, err := d.elements.Get(0)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// Peek returns the first element in the Queue.
// Returns -1 if Queue is un-initialized or has no more elements to remove(empty Queue).
func (d *DllQueue[T]) Peek() (T, error) {
	var ret T
	if d == nil {
		return ret, errors.UninitializedError
	}
	if d.elements.Size() == 0 || d.elements == nil {
		return ret, errors.UninitializedError
	}
	ret, err := d.elements.Get(0)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// Clear the queue.
func (d *DllQueue[T]) Clear() {
	d.elements.Clear()
}

// String implements stringer interface.
func (d *DllQueue[T]) String() string {
	return d.elements.String()
}
