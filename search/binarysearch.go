package search

import "github.com/maverick6912/dsa_go/errors"

// Represents the direction of sort.
type SortDirection int

const (
	Ascending SortDirection = iota
	Descending
)

// BinarySearch searches a sorted(ascending or descending) array for given key using the comparator passed.
// If key is found it returns the index at which the key was found and error is nil.
// If key is not found it returns -1 and NotFound error.
// It also returns -1 and NoElements error if iterable passed has no elements.
func BinarySearch[T any](i []T, k T, sortDir SortDirection, cmp func(T, T) int) (int, error) {
	if len(i) == 0 {
		return -1, errors.NoElements
	}
	if sortDir == Ascending {
		return binarySearchForAscending(i, k, cmp)
	} else {
		return binarySearchForDescending(i, k, cmp)
	}
}

// binarySearchForAscending performs the actual search in ascending sorted array.
// If key is not found it returns -1 and NotFound error.
func binarySearchForAscending[T any](i []T, k T, cmp func(T, T) int) (int, error) {
	if !isAscendingSorted(i, cmp) {
		return -1, errors.ErrImproperSort
	}
	low, high := 0, len(i)-1
	for low <= high {
		mid := low + (high-low)/2
		// {1,2,3,4,5}
		if cmp(i[mid], k) > 0 {
			high = mid - 1
		} else if cmp(i[mid], k) < 0 {
			low = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, errors.NotFound
}

// binarySearchForAscending performs the actual search in descending sorted array.
// If key is not found it returns -1 and NotFound error.
func binarySearchForDescending[T any](i []T, k T, cmp func(T, T) int) (int, error) {
	if !isDescendingSorted(i, cmp) {
		return -1, errors.ErrImproperSort
	}
	low, high := 0, len(i)-1
	for low <= high {
		mid := low + (high-low)/2
		// {5,4,3,2,1}
		if cmp(i[mid], k) > 0 {
			// i[mid] > k
			low = mid + 1
		} else if cmp(i[mid], k) < 0 {
			// i[mid] < k
			high = mid - 1
		} else {
			return mid, nil
		}
	}
	return -1, errors.NotFound
}

// isAscendingSorted checks if array is sorted in ascending.
func isAscendingSorted[T any](it []T, cmp func(T, T) int) bool {
	ret := true
	for i := 0; i < len(it)-1; i++ {
		if cmp(it[i], it[i+1]) > 0 {
			ret = false
		}
	}
	return ret
}

// isAscendingSorted checks if array is sorted in descending.
func isDescendingSorted[T any](it []T, cmp func(T, T) int) bool {
	ret := true
	for i := 0; i < len(it)-1; i++ {
		if cmp(it[i], it[i+1]) < 0 {
			ret = false
		}
	}
	return ret
}
