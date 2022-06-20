package heap

import "github.com/maverick6912/dsa_go/errors"

type Heap[T any] struct {
	cmp   func(T, T) int
	elems []T
	len   int
}

type MaxHeap[T any] Heap[T]

// New initializes a heap and returns it.
func (m *MaxHeap[T]) New(cmp func(T, T) int) *MaxHeap[T] {
	return &MaxHeap[T]{cmp: cmp, elems: make([]T, 0)}
}

// Insert a value to heap.
// returns uninitialized error if heap is not initialized.
func (m *MaxHeap[T]) Insert(val T) {
	m.elems = append(m.elems, val)
	m.maxHeapifyUp(len(m.elems) - 1)
	m.len += 1
}

//maxHeapifyUp heapifies from bottom to top.
func (m *MaxHeap[T]) maxHeapifyUp(index int) {
	for m.cmp(m.elems[parent(index)], m.elems[index]) < 0 {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

// Extract returns and deletes the largest element from heap.
func (m *MaxHeap[T]) Extract() (T, error) {
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
	m.maxHeapifyDown(0)
	m.len -= 1
	return extracted, nil
}

//maxHeapifyDown heapifies from top to bottom.
func (m *MaxHeap[T]) maxHeapifyDown(index int) {
	lastIndex := len(m.elems) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.cmp(m.elems[l], m.elems[r]) > 0 {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.cmp(m.elems[index], m.elems[childToCompare]) < 0 {
			m.swap(index, childToCompare)
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
func (m *MaxHeap[T]) swap(index1, index2 int) {
	m.elems[index1], m.elems[index2] = m.elems[index2], m.elems[index1]
}

// Length returns current length of heap.
func (m *MaxHeap[T]) Length() int {
	return m.len
}

// Values returns the elements in the heap.
func (m Heap[T]) Values() []T {
	return m.elems
}
