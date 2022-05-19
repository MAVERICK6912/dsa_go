package main

/*
Approach:
	Input arr[] = [1, 2, 3, 4, 5, 6, 7], d = 2, n =7
		1) Store the first d elements in a temp array
			temp[] = [1, 2]
		2) Shift rest of the arr[]
			arr[] = [3, 4, 5, 6, 7, 6, 7]
		3) Store back the d elements
			arr[] = [3, 4, 5, 6, 7, 1, 2]
*/

// func main() {
// 	fmt.Print(rotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 7, 7))
// }

func rotateArray(arr []int, d, n int) []int {
	if d == len(arr) {
		return arr
	}
	var tempArr []int
	for i := 0; i < d; i++ {
		tempArr = append(tempArr, arr[i])
	}
	for i := 0; i < len(arr)-d; i++ {
		arr[i] = arr[i+d]
	}
	for i := 0; i < d; i++ {
		arr[(len(arr)+i)-(d)] = tempArr[i]
	}
	return arr
}
