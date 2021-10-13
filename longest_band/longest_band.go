package main

import "fmt"

/*
	Problem: In this problem we need to find consecutive numbers(with a difference of 1) from a given array/sliceand return the largest sequence that can be formed with these consecutive numbers.
	Problem link: https://leetcode.com/problems/longest-consecutive-sequence/
*/

/*
Approach:
	1. Put the numbers in a hashmap or a dictionary(this will take O(N) time)
	2. Loop over all the elements in the array/slice.
	3. Check if there's a parent element(current_element-1) for the current element.
		a. If true then don't do anything.
		b. If false then try to find the next band element and keep finding bands for elements.
*/
func findLongestBand(input []int) int {
	elementMap := make(map[int]bool, 0)
	for _, v := range input {
		// save each element in a map/dictionary
		elementMap[v] = true
	}
	largestCount := 0
	/*
		 TIL: Iterating over a map is much faster than a slice in this case.
		 In this case this is possible as:
			When iterating over the slice then accessing the map, a larger block of memory will need to make its way back and forth between DRAM and the CPU cache than if you only use the map
	*/
	for key := range elementMap {
		parent := key - 1
		// Check if parent element is present in the map.
		if _, present := elementMap[parent]; !present {
			nextNumber := key + 1
			// since we're already on the element we initialize count as 1.
			count := 1
			// loop until we keep finding next_number elements in map.
			for elementMap[nextNumber] {
				nextNumber++
				count++
			}
			if count > largestCount {
				largestCount = count
			}
		}
	}
	return largestCount
}

func main() {
	fmt.Println("Largest band: ", findLongestBand([]int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6}))
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
