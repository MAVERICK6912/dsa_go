package main

import "fmt"

// In a pairs problem we need to find a pairs of distinct elements in an array which add up to a target number.
// 1st approach(brute force): O(N^2)
// 2nd approach(sorting and searching): O(NlogN)+O(NlogN)=>O(NlogN) first sort all elems and then for every elem do a search for corresponding pair.
// 3rd approach(using a map/hashset): O(N)

func findPair(arr []int, targetSum int) [][]int {
	elemMap := make(map[int]bool, 0)
	ret := make([][]int, 0)
	for _, val := range arr {
		elemToFind := targetSum - val
		if elemMap[elemToFind] {
			ret = append(ret, []int{val, elemToFind})
		}
		elemMap[val] = true
	}
	return ret
}

func main() {
	l := []int{10, 5, 2, 3, -6, 9, 11, 2, 1}
	for _, v := range findPair(l, 4) {
		fmt.Println(v)
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
