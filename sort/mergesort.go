package sort

import "github.com/maverick6912/dsa_go/errors"

// MergeSort sorts the array as per the sortDirection passed.
// returns NoElements error if passed iterable is empty.
func MergeSort[T any](i []T, sortDir SortDirection, cmp func(T, T) int) ([]T, error) {
	if len(i) == 0 {
		return nil, errors.NoElements
	}
	l, r := 0, len(i)-1

	return divideAndMerge(&i, l, r, cmp, sortDir), nil
}

// divideAndMerge divides the slice recursively and the calls merge on each sub-slice so formed.
func divideAndMerge[T any](i *[]T, l, r int, cmp func(T, T) int, sortDir SortDirection) []T {
	if l < r {
		m := l + (r-l)/2
		// divide left half
		divideAndMerge(i, l, m, cmp, sortDir)
		// divide right half
		divideAndMerge(i, m+1, r, cmp, sortDir)
		// merge the divided array/slice
		if sortDir == Ascending {
			t := mergeAscending(*i, l, r, m, cmp)
			i = &t
		} else {
			t := mergeDescending(*i, l, r, m, cmp)
			i = &t
		}
	}
	return *i
}

// mergeSortAscending sorts the sub slices while merging them to the original slice in ascending order.
func mergeAscending[T any](it []T, l, r, m int, cmp func(T, T) int) []T {
	// size of temp slices
	subSliceOneSize, subSliceTwoSize := m-l+1, r-m
	// initializing temp slices
	leftTempSlice, rightTempSlice := make([]T, 0), make([]T, 0)

	// copying data to temp slices
	for i := 0; i < subSliceOneSize; i++ {
		leftTempSlice = append(leftTempSlice, it[l+i])
	}
	for i := 0; i < subSliceTwoSize; i++ {
		rightTempSlice = append(rightTempSlice, it[m+i+1])
	}

	// initial indexes of temp slice
	iSubSliceOne, iSubAliceTwo := 0, 0
	// initial index of merged sub-slice
	k := l

	// looping till index of subSliceOne<size of subSliceOne and index of subSliceTwo<size of subSliceTwo
	for iSubSliceOne < subSliceOneSize && iSubAliceTwo < subSliceTwoSize {
		// comparing if leftTempSlice[iSubSliceOne]<=rightTempSlice[iSubAliceTwo]
		if cmp(leftTempSlice[iSubSliceOne], rightTempSlice[iSubAliceTwo]) < 0 || cmp(leftTempSlice[iSubSliceOne], rightTempSlice[iSubAliceTwo]) == 0 {
			// if above condition is true then swap originalSlice(it)[currentIndexOfMergedArray] with leftTempSlice[iSubSliceOne] in place.
			it[k] = leftTempSlice[iSubSliceOne]
			iSubSliceOne += 1
		} else {
			// if above condition is false then swap originalSlice(it)[currentIndexOfMergedArray] with rightTempSlice[iSubAliceTwo] in place.
			it[k] = rightTempSlice[iSubAliceTwo]
			iSubAliceTwo += 1
		}
		k += 1
	}
	// putting rest of the elements of tempLeftSubSlice back in place.
	for iSubSliceOne < subSliceOneSize {
		it[k] = leftTempSlice[iSubSliceOne]
		iSubSliceOne += 1
		k += 1
	}
	// putting rest of the elements of tempRightSubSlice back in place.
	for iSubAliceTwo < subSliceTwoSize {
		it[k] = rightTempSlice[iSubAliceTwo]
		iSubAliceTwo += 1
		k += 1
	}
	return it
}

// mergeSortDescending sorts the sub slices while merging them to the original slice in descending order.
func mergeDescending[T any](it []T, l, r, m int, cmp func(T, T) int) []T {
	// size of temp slices
	subSliceOneSize, subSliceTwoSize := m-l+1, r-m
	// initializing temp slices
	leftTempSlice, rightTempSlice := make([]T, 0), make([]T, 0)

	// copying data to temp slices
	for i := 0; i < subSliceOneSize; i++ {
		leftTempSlice = append(leftTempSlice, it[l+i])
	}
	for i := 0; i < subSliceTwoSize; i++ {
		rightTempSlice = append(rightTempSlice, it[m+i+1])
	}

	// initial indexes of temp slice
	iSubSliceOne, iSubAliceTwo := 0, 0
	// initial index of merged sub-slice
	k := l
	// looping till index of subSliceOne<size of subSliceOne and index of subSliceTwo<size of subSliceTwo
	for iSubSliceOne < subSliceOneSize && iSubAliceTwo < subSliceTwoSize {
		// comparing if leftTempSlice[iSubSliceOne]<=rightTempSlice[iSubAliceTwo]
		if cmp(leftTempSlice[iSubSliceOne], rightTempSlice[iSubAliceTwo]) < 0 || cmp(leftTempSlice[iSubSliceOne], rightTempSlice[iSubAliceTwo]) == 0 {
			// if above condition is true then swap originalSlice(it)[currentIndexOfMergedArray] with rightTempSlice[iSubAliceTwo] in place.
			it[k] = rightTempSlice[iSubAliceTwo]
			iSubAliceTwo += 1
		} else {
			// if above condition is false then swap originalSlice(it)[currentIndexOfMergedArray] with leftTempSlice[iSubSliceOne] in place.
			it[k] = leftTempSlice[iSubSliceOne]
			iSubSliceOne += 1
		}
		k += 1
	}
	// putting rest of the elements of tempLeftSubSlice back in place.
	for iSubSliceOne < subSliceOneSize {
		it[k] = leftTempSlice[iSubSliceOne]
		iSubSliceOne += 1
		k += 1
	}
	// putting rest of the elements of tempRightSubSlice back in place.
	for iSubAliceTwo < subSliceTwoSize {
		it[k] = rightTempSlice[iSubAliceTwo]
		iSubAliceTwo += 1
		k += 1
	}
	return it
}
