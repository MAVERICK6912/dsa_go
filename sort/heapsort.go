package sort

import (
	"github.com/maverick6912/dsa_go/errors"
	heap "github.com/maverick6912/dsa_go/heap/basics"
)

// HeapSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty
// and HeapExtract error if there waas some issue while extracting from heap.
func HeapSort[T any](it []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(it) == 0 {
		return nil, errors.NoElements
	}
	if sortDir == Ascending {
		return sortAscending(it, cmp)
	} else {
		return sortDescending(it, cmp)
	}
}

// sortAscending sorts the slice in ascending order by first making a MinHeap and the extracting from it.
func sortAscending[T any](it []T, cmp func(T, T) int) ([]T, error) {
	var minHeap *heap.MinHeap[T]
	minHeap = minHeap.New(cmp)

	for _, v := range it {
		minHeap.Insert(v)
	}
	for i := range it {
		val, err := minHeap.Extract()
		if err != nil {
			return nil, errors.ErrHeapExtract
		}
		it[i] = val
	}
	return it, nil
}

// sortDescending sorts the slice in descending order by first making a MaxHeap and the extracting from it.
func sortDescending[T any](it []T, cmp func(T, T) int) ([]T, error) {
	var maxHeap *heap.MaxHeap[T]
	maxHeap = maxHeap.New(cmp)

	for _, v := range it {
		maxHeap.Insert(v)
	}
	for i := range it {
		val, err := maxHeap.Extract()
		if err != nil {
			return nil, errors.ErrHeapExtract
		}
		it[i] = val
	}
	return it, nil
}
