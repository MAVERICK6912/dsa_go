package main

import (
	"fmt"
	"math"
)

/*
	Problem: given an array findout a sub-array such that sorting that sub-array sorts the whole array.
	For example: {1, 2, 3, 4, 5, 8, 6, 7, 9, 10, 11}, given this array.
		index:	  0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
		We can see that element from indices 5 to 7 are not sorted. Hence if we sort them then the whole array will be sorted.
*/
/*
	First approach: Time complexity O(nlog(n))
	1. Sort the given array and store as a separate array. Time coplexity for this op: O(nlog(n))
	2. Loop over input array and compare each element with the sorted array. Find the elements which do not match. Time coplexity for this op: O(n)

	Total Time Complexity: O(nlog(n))+O(n)=O(nlog(n))

	Second approach: Time complexity O(n)
	1. An ascending sorted element will follow: input[i-1] < input[i] < input[i+1]. Where 'i' is the current element in the array.
	2. Find the smallest and largest element that are out of order using the above explanation. Time coplexity for this op: O(n)
	3. Find the index of smallest and largest element out of order. Time coplexity for this op: O(n)

	Total time complexity: O(n)+O(n)=O(n)

*/

// Problem link: https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

type response struct {
	startIndex int
	endIndex   int
}

func outOfOrder(input []int, index int) bool {
	if index == 0 {
		return input[index] > input[index+1]
	}
	if index == len(input)-1 {
		return input[index] < input[index-1]
	}
	return input[index] > input[index+1] || input[index] < input[index-1]
}

func subArraySort(input []int) response {
	smallestOutOfOrder, largestOutOfORder := math.MaxInt, math.MinInt

	for index, val := range input {
		if outOfOrder(input, index) {
			smallestOutOfOrder = smaller(smallestOutOfOrder, val)
			largestOutOfORder = bigger(largestOutOfORder, val)
		}
	}
	if smallestOutOfOrder == math.MaxInt {
		return response{startIndex: -1, endIndex: -1}
	}

	left := 0
	for smallestOutOfOrder >= input[left] {
		left += 1
	}
	right := len(input) - 1
	for largestOutOfORder <= input[right] {
		right -= 1
	}
	return response{startIndex: left, endIndex: right}
}
func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(subArraySort([]int{1, 2, 3, 4, 5, 8, 6, 7, 9, 10, 11}))
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
