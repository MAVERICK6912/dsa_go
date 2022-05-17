package main

import "fmt"

// problem link: https://leetcode.com/problems/richest-customer-wealth/

func main() {
	fmt.Print(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		totWealth := 0
		for _, wealthInBank := range customer {
			totWealth += wealthInBank
		}
		if maxWealth < totWealth {
			maxWealth = totWealth
		}
	}
	return maxWealth
}
