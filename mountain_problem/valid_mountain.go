package main

// An array input is a valid mountain if it is strictly increasing till peak and the strictly decreasing till the end of array.
// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// 	arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// 	arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// problem link: https://leetcode.com/problems/valid-mountain-array/

func main() {
	// put input code or hard code input
}

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	} else {
		largest := 0
		for i := 1; i <= n-2; {
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				count := 1
				// count backwards
				j := i
				for j >= 1 && arr[j] > arr[j-1] {
					j--
					count++
				}
				// count backwards
				for i <= n-2 && arr[i] > arr[i+1] {
					i++
					count++
				}
				largest = max(largest, count)
			} else {
				i++
			}
		}
		if largest != len(arr) {
			return false
		} else {
			return true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
