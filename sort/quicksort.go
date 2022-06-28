package sort

import (
	"github.com/maverick6912/dsa_go/errors"
)

func QuickSort[T any](it []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(it) == 0 {
		return nil, errors.NoElements
	}
	low, high := 0, len(it)-1
	return partitionAnQuickSort(it, low, high, sortDir, cmp), nil
}

func partitionAnQuickSort[T any](it []T, low, high int, sortDir SortDirection, cmp func(T, T) int) []T {
	if low < high {
		var partitionIndex int
		if sortDir == Ascending {
			partitionIndex = partitionAscending(it, low, high, cmp)
		} else {
			partitionIndex = partitionDescending(it, low, high, cmp)
		}

		partitionAnQuickSort(it, low, partitionIndex-1, sortDir, cmp)
		partitionAnQuickSort(it, partitionIndex+1, high, sortDir, cmp)
	}
	return it
}

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

// func swap[T any](a, b *T) {
// 	t := a
// 	a = b
// 	b = t
// }
