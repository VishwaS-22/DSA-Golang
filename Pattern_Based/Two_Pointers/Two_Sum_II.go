package main

func two_sum(arr []int, t int) []int {
	l, r := 0, len(arr)-1

	for l < r {
		s := arr[l] + arr[r]

		if s == t {
			return []int{l + 1, r + 1}
		} else if s > t {
			r--
		} else {
			l++
		}

	}
	return []int{}
}
