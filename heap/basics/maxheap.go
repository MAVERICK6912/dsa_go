package heap

import "github.com/maverick6912/dsa_go/errors"

type Heap[T any] struct {
	cmp   func(T, T) int
	elems []T
	len   int
}

// New initializes a heap and returns it.
func (h *Heap[T]) New(cmp func(T, T) int) *Heap[T] {
	return &Heap[T]{cmp: cmp, elems: make([]T, 0)}
}

// Insert a value to heap.
// returns uninitialized error if heap is not initialized.
func (h *Heap[T]) Insert(val T) {
	h.elems = append(h.elems, val)
	h.maxHeapifyUp(len(h.elems) - 1)
	h.len += 1
}

//maxHeapifyUp heapifies from bottom to top.
func (h *Heap[T]) maxHeapifyUp(index int) {
	for h.cmp(h.elems[parent(index)], h.elems[index]) < 0 {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Extract returns and deletes the largest element from heap.
func (h *Heap[T]) Extract() (T, error) {
	var extracted T
	if h == nil {
		return extracted, errors.UninitializedError
	}
	if h.Length() == 0 {
		return extracted, errors.NoElements
	}
	extracted = h.elems[0]
	lastIndex := len(h.elems) - 1
	h.elems[0] = h.elems[lastIndex]
	h.elems = h.elems[:lastIndex]
	h.maxHeapifyDown(0)
	h.len -= 1
	return extracted, nil
}

//maxHeapifyDown heapifies from top to bottom.
func (h *Heap[T]) maxHeapifyDown(index int) {
	lastIndex := len(h.elems) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.cmp(h.elems[l], h.elems[r]) > 0 {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.cmp(h.elems[index], h.elems[childToCompare]) < 0 {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent returns the parent of the index passed.
func parent(index int) int {
	return (index - 1) / 2
}

// left returns the index of left child of parent.
func left(parentIndex int) int {
	return 2*parentIndex + 1
}

// right returns the index of right child of parent.
func right(parentIndex int) int {
	return 2*parentIndex + 2
}

// swap value at the passed indices.
func (h *Heap[T]) swap(index1, index2 int) {
	h.elems[index1], h.elems[index2] = h.elems[index2], h.elems[index1]
}

// Length returns current length of heap.
func (h *Heap[T]) Length() int {
	return h.len
}

// Values returns the elements in the heap.
func (h Heap[T]) Values() []T {
	return h.elems
}
