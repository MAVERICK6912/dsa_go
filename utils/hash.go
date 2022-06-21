package utils

type Hasher[K any] func(K, int) int

func StringHash[K string](key K, size int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return (sum % len(key)) % size
}

func IntHash[K int](key K, size int) int {
	sum, t, c := 0, key, 0
	for t != 0 {
		sum += int(t) % 10
		t /= 10
		c += 1
	}
	return (sum % c) % size
}
