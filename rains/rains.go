package main

import "fmt"

/*
Problem: given an array like: {0,1,0,2,1,0,1,3,2,1,2,1}
    ↓ ↓  ↓   ↓  ↓
              |
     |        || |
__|__|__|__|__|||||
0 1 020 10 10 32121

This is kind-of how the elevation map would like for the above array.
For the given array input we need to find out how much water will be trapped in it.

Take aways:
1. Amount of water trapped will always be dependent on the height of lowest wall in the container.

Problem link: https://leetcode.com/problems/trapping-rain-water
*/

/*
Approach:
1. Scan from right and make a new slice with highest element till another highest is found from right to left.
	Example:
	input:        {0,1,0,2,1,0,1,3,2,1,2,1}
	rightHighest: {3,3,3,3,3,3,3,3,2,2,2,1}
2. Scan from left and make a new slice with highest element till another highest is found from left to right.
	Example:
	input:        {0,1,0,2,1,0,1,3,2,1,2,1}
	rightHighest: {0,1,1,2,2,2,2,3,3,3,3,3}
3. Loop over the input array and use this equation: min(rightHighest[i],leftHighest[i])-input[i]
	Example:
	-> min(rightHighest[0],leftHighest[0])-input[0]:
	-> min(0,3)-0
	-> 0-0=0

	-> min(rightHighest[1],leftHighest[1])-input[1]:
	-> min(3,1)-1
	-> 1-1=0

	-> min(rightHighest[2],leftHighest[2])-input[2]:
	-> min(3,1)-0
	-> 1-0=1
*/

func findAmountOfTrappedWater(input []int) int {
	leftHigh, rightHigh := make([]int, len(input)), make([]int, len(input))
	highest := 0
	for index, val := range input {
		if highest <= val {
			highest = val
		}
		leftHigh[index] = highest
	}
	highest = 0
	for i := len(input) - 1; i >= 0; i-- {
		if highest < input[i] {
			highest = input[i]
		}
		rightHigh[i] = highest
	}
	waterTrapped := 0
	for index, val := range input {
		waterTrapped += findMin(leftHigh[index], rightHigh[index]) - val
	}
	return waterTrapped
}

func findMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Print("Water trapped: ", findAmountOfTrappedWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

/*
                                 .__        __     ________________  ____________
  _____ _____ ___  __ ___________|__| ____ |  | __/  _____/   __   \/_   \_____  \
 /     \\__  \\  \/ // __ \_  __ \  |/ ___\|  |/ /   __  \\____    / |   |/  ____/
|  Y Y  \/ __ \\   /\  ___/|  | \/  \  \___|    <\  |__\  \  /    /  |   /       \
|__|_|  (____  /\_/  \___  >__|  |__|\___  >__|_ \\_____  / /____/   |___\_______ \
      \/     \/          \/              \/     \/      \/                       \/

*/
