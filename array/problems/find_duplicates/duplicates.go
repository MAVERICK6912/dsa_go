package main

import "fmt"

func main() {

	fmt.Print(FindDuplicates([]int{23, 34, 54, 34, 23, 43, 54, 67}))

}

func FindDuplicates(arr []int) []int {
	elemCountMap := make(map[int]int)
	for _, v := range arr {
		if _, ok := elemCountMap[v]; !ok {
			elemCountMap[v] = 1
		} else {
			elemCountMap[v] += 1
		}
	}
	var retSlice []int
	for k, v := range elemCountMap {
		if v > 1 {
			retSlice = append(retSlice, k)
		}
	}
	if len(retSlice) > 0 {
		return retSlice
	} else {
		retSlice = append(retSlice, -1)
		return retSlice
	}
}
