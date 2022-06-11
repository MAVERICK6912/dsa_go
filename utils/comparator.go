package utils

type Comparator[T any] func(a, b T) int

func IntComparator(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func StringComparator(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
