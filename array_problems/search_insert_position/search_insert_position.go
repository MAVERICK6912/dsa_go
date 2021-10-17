package main

import "fmt"

/*
	Problem: Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

	Approach: Time complexity- O(n)
	1. Loop over the array.
	2. If target is found return index.
	3. Else: find the position where the target would fit in:
		a. If target is less than input[0] then return 0.
		b. If target is greater than input[len(input)-1] then return len(input).
		c. If target is greater than input[index-1] and target is less than input[index]. Return current index.
		d. If target is greater than input[index] and target is less than input[index+1]. Return current index+1.
*/

// Problem link: https://leetcode.com/problems/search-insert-position/

func searchInsert(input []int, target int) int {
	for index, val := range input {
		if val == target {
			return index
		}
	}
	for index, val := range input {
		if index == 0 {
			if target < val {
				return 0
			}
		}
		if index == (len(input) - 1) {
			if target > val {
				return len(input)
			}
		}
		if index > 0 && index < len(input)-1 {
			if target > input[index-1] && target < input[index] {
				return index
			}
			if target > input[index] && target < input[index+1] {
				return index + 1
			}
		}
	}
	// Handling corner cases
	if len(input) < 3 {
		if target < input[0] {
			return 0
		} else if target > input[0] && target < input[1] {
			return 1
		}
		return len(input)
	}
	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
