package bella

import "cmp"

func compare[T cmp.Ordered](a, b T) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	}
	return 1
}

func zero[T any]() T {
	var t T
	return t
}
