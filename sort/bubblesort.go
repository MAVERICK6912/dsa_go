package sort

import "github.com/maverick6912/dsa_go/errors"

// BubbleSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty.
func BubbleSort[T any](i []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(i) == 0 {
		return nil, errors.NoElements
	}
	if sortDir == Ascending {
		return bubbleSortAscending(i, cmp), nil
	} else {
		return bubbleSortDescending(i, cmp), nil
	}
}

// bubbleSortAscending sorts the array in ascending order.
func bubbleSortAscending[T any](i []T, cmp func(T, T) int) []T {
	for m := range i {
		swapped := false
		for n := 0; n < len(i)-m-1; n++ {
			if cmp(i[n], i[n+1]) > 0 {
				temp := i[n]
				i[n] = i[n+1]
				i[n+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return i
}

// bubbleSortDescending sorts the array in descending order.
func bubbleSortDescending[T any](i []T, cmp func(T, T) int) []T {
	for m := range i {
		swapped := false
		for n := 0; n < len(i)-m-1; n++ {
			if cmp(i[n], i[n+1]) < 0 {
				temp := i[n]
				i[n] = i[n+1]
				i[n+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return i
}
