package utils

type Comparator func(a, b int) int

func IntComparator(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
