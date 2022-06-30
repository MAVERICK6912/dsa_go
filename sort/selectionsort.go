package sort

import (
	"github.com/maverick6912/dsa_go/errors"
)

// SelectionSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty.
func SelectionSort[T any](i []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(i) == 0 {
		return nil, errors.NoElements
	}
	if sortDir == Ascending {
		return selectionSortAscending(i, cmp), nil
	} else {
		return selectionSortDescending(i, cmp), nil
	}
}

// selectionSortAscending sorts the array in ascending order.
func selectionSortAscending[T any](it []T, cmp func(T, T) int) []T {
	for m := range it {
		minIndex := m
		for n := m + 1; n < len(it); n++ {
			if cmp(it[n], it[minIndex]) < 0 {
				minIndex = n
			}
		}
		temp := it[minIndex]
		it[minIndex] = it[m]
		it[m] = temp
	}
	return it
}

// selectionSortDescending sorts the array in descending order.
func selectionSortDescending[T any](it []T, cmp func(T, T) int) []T {
	for m := range it {
		maxIndex := m
		for n := m + 1; n < len(it); n++ {
			if cmp(it[n], it[maxIndex]) > 0 {
				maxIndex = n
			}
		}
		temp := it[maxIndex]
		it[maxIndex] = it[m]
		it[m] = temp
	}
	return it
}
