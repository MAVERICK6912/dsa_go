package search

import "github.com/maverick6912/dsa_go/errors"

/*
	This implementation of linear search uses two pointer approach. It works as follows:
	- left,right:=0,len(arr)-1
	- loop till left<=right:
		- if arr[left]==k:
			- return true(or we can return the element)
		- if arr[right]==k:
			- return true(or we can return the element)
		right-=1
		left+=1
	return false
*/
// LinearSearch searches for the given key using the comparator passed.
// If key is found it returns the index at which the key was found and error is nil.
// If key is not found it returns -1 and NotFound error.
// It also returns -1 and NoElements error if iterable passed has no elements.
func LinearSearch[T any](i []T, k T, cmp func(T, T) int) (int, error) {
	if len(i) == 0 {
		return -1, errors.NoElements
	}
	left, right := 0, len(i)-1
	for left <= right {
		if cmp(i[left], k) == 0 {
			return left, nil
		}
		if cmp(i[right], k) == 0 {
			return right, nil
		}
		left += 1
		right -= 1
	}
	return -1, errors.NotFound
}
