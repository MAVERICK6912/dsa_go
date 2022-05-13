package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("a  a"))
}

func lengthOfLastWord(input string) int {
	count := 0

	for i := len(input) - 1; i >= 0; i-- {
		if count == 0 && input[i] == ' ' {
			continue
		} else {
			if input[i] == ' ' {
				break
			}
			count += 1
		}

	}
	return count
}
