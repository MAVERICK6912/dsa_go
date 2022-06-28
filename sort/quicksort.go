package sort

import (
	"github.com/maverick6912/dsa_go/errors"
)

// QuickSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty.
func QuickSort[T any](it []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(it) == 0 {
		return nil, errors.NoElements
	}
	low, high := 0, len(it)-1
	return partitionAndQuickSort(it, low, high, sortDir, cmp), nil
}

// partitionAnQuickSort partitions and sorts the slice recursively.
func partitionAndQuickSort[T any](it []T, low, high int, sortDir SortDirection, cmp func(T, T) int) []T {
	if low < high {
		var partitionIndex int
		if sortDir == Ascending {
			partitionIndex = partitionAscending(it, low, high, cmp)
		} else {
			partitionIndex = partitionDescending(it, low, high, cmp)
		}

		partitionAndQuickSort(it, low, partitionIndex-1, sortDir, cmp)
		partitionAndQuickSort(it, partitionIndex+1, high, sortDir, cmp)
	}
	return it
}

// partitionAscending partitions and sorts the slice in ascending order
func partitionAscending[T any](it []T, low, high int, cmp func(T, T) int) int {
	pivot, pRight := it[high], low-1
	for i := low; i < high; i++ {
		if cmp(it[i], pivot) < 0 {
			pRight += 1
			t := it[pRight]
			it[pRight] = it[i]
			it[i] = t
		}
	}
	t := it[pRight+1]
	it[pRight+1] = it[high]
	it[high] = t

	return pRight + 1
}

// partitionDescending partitions and sorts the slice in descending order
func partitionDescending[T any](it []T, low, high int, cmp func(T, T) int) int {
	pivot, pRight := it[high], low-1
	for i := low; i < high; i++ {
		if cmp(it[i], pivot) > 0 {
			pRight += 1
			t := it[pRight]
			it[pRight] = it[i]
			it[i] = t
		}
	}
	t := it[pRight+1]
	it[pRight+1] = it[high]
	it[high] = t

	return pRight + 1
}
