package main

import (
	"fmt"
	"math"
)

// In this problem we need to identify the peak in a given input and return its width.
// example: [5,6,1,2,3,4,5,4,3,2,0,1,2,3,-2,4]
// here we can see that the peak is at index 6 and the width of this mountain is 9 as there are 9 numbers forming this mountain

/* Approach:
1. Identify peaks. A peak is a number with lower numbers as neigbors. example: [1,3,2] 3 is the peak.
2. then iterate back to count left slope.
3. since while finding peaks we're already moving forward so we can count right slope in that pass.
*/
func main() {
	fmt.Println(highestMountain([]int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4}))
}

func highestMountain(input []int) int {
	largest := 0
	if len(input) < 3 {
		return 0
	} else {
		for i := 1; i <= len(input)-2; {
			// check if input[i] is a peak
			if input[i] > input[i-1] && input[i] > input[i+1] {
				count := 1
				// count backwards
				j := i
				for j >= 1 && input[j] > input[j-1] {
					j--
					count++
				}
				// count forwards
				for i <= len(input)-2 && input[i] > input[i+1] {
					i++
					count++
				}
				largest = int(math.Max(float64(largest), float64(count)))
			} else {
				i++
			}
		}
		return largest
	}
}
