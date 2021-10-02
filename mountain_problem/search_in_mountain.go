package main

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	peakIndex := findPeakElementsIndex(mountainArr)

	var index int = -1
	// search in first half of the array
	index = findElementinSortedArray(peakIndex, target, mountainArr)
	if index == -1 {
		// second half of the array but this array is sorted in descending order
		index = findElementInReversedArray(peakIndex, target, mountainArr)
	}

	return index
}

func findPeakElementsIndex(mountainArr *MountainArray) int {

	var left, right, mid int
	right = mountainArr.length() - 1

	for left < right {
		mid = left + (right-left)/2

		midElement := mountainArr.get(mid)
		nextElement := mountainArr.get(mid + 1)

		if midElement < nextElement {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findElementinSortedArray(peakIndex, target int, mountainArr *MountainArray) int {

	var left, right, mid int
	right = peakIndex

	for left <= right {
		mid = left + (right-left)/2

		midElement := mountainArr.get(mid)
		if midElement == target {
			return mid
		} else if midElement < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findElementInReversedArray(peakIndex, target int, mountainArr *MountainArray) int {

	var left, right, mid int
	left = peakIndex + 1
	right = mountainArr.length() - 1

	for left <= right {
		mid = left + (right-left)/2

		midElement := mountainArr.get(mid)
		if midElement == target {
			return mid
		} else if midElement < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
