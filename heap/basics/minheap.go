package heap

import "github.com/maverick6912/dsa_go/errors"

type MinHeap[T any] Heap[T]

// New initializes a heap and returns it.
func (m *MinHeap[T]) New(cmp func(T, T) int) *MinHeap[T] {
	return &MinHeap[T]{cmp: cmp, elems: make([]T, 0)}
}

// Insert a value to heap.
// returns uninitialized error if heap is not initialized.
func (m *MinHeap[T]) Insert(val T) {
	m.elems = append(m.elems, val)
	m.minHeapifyUp(len(m.elems) - 1)
	m.len += 1
}

//maxHeapifyUp heapifies from bottom to top.
func (m *MinHeap[T]) minHeapifyUp(index int) {
	for m.cmp(m.elems[parent(index)], m.elems[index]) > 0 {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

// Extract returns and deletes the largest element from heap.
func (m *MinHeap[T]) Extract() (T, error) {
	var extracted T
	if m == nil {
		return extracted, errors.UninitializedError
	}
	if m.Length() == 0 {
		return extracted, errors.NoElements
	}
	extracted = m.elems[0]
	lastIndex := len(m.elems) - 1
	m.elems[0] = m.elems[lastIndex]
	m.elems = m.elems[:lastIndex]
	m.minHeapifyDown(0)
	m.len -= 1
	return extracted, nil
}

//maxHeapifyDown heapifies from top to bottom.
func (m *MinHeap[T]) minHeapifyDown(index int) {
	lastIndex := len(m.elems) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.cmp(m.elems[l], m.elems[r]) < 0 {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.cmp(m.elems[index], m.elems[childToCompare]) > 0 {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// swap value at the passed indices.
func (m *MinHeap[T]) swap(index1, index2 int) {
	m.elems[index1], m.elems[index2] = m.elems[index2], m.elems[index1]
}

// Length returns current length of heap.
func (m *MinHeap[T]) Length() int {
	return m.len
}
