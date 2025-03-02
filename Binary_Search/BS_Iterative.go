package main

func BS_Iterative(t int, arr []int) int {
	n := len(arr)
	l, h := 0, n-1

	for l <= h {
		m := (l + h) / 2
		if arr[m] == t {
			return m
		} else if t > arr[m] {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}
