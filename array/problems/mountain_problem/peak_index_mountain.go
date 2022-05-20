package main

// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// 	 arr[0] < arr[1] < ... arr[i-1] < arr[i]
// 	 arr[i] > arr[i+1] > ... > arr[arr.length - 1]

// problem link: https://leetcode.com/problems/peak-index-in-a-mountain-array/

// func main() {
// 	// put input code or hard code input
// }

func peakIndexInMountainArray(arr []int) int {
	if len(arr) < 3 {
		return 0
	} else {
		for i := 1; i <= len(arr)-2; {
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				return i
			} else {
				i++
			}
		}
	}
	return 0
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
