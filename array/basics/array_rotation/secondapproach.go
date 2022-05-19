package main

/*
	Approach:
		rotateArrayRecur(arr[], d, currIter)
		start
		temp=arr[0]
		For i = 0 to i < d
			Left rotate all elements of arr[] by one
		swap(arr[last],temp)
		if currIter == d
			return
		currIter+=1
		return rotateArrayRecur(arr[], d, currIter)
		end
*/

// func main() {
// 	fmt.Print(rotateArrayRecur([]int{1, 2, 3, 4, 5, 6, 7}, 5, 1))

// }

func rotateArrayRecur(arr []int, d, currIter int) ([]int, int, int) {
	if len(arr) == d {
		return arr, d, currIter
	}
	temp := arr[0]
	for i := 0; i < len(arr)-currIter; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-currIter] = temp
	if currIter == d {
		return arr, d, currIter
	}
	currIter += 1
	return rotateArrayRecur(arr, d, currIter)
}
