package main

// floor is finding the largest num <= x
func floor(s, e int, arr []int) int {
	v := -1
	l, h := 0, s-1
	for l <= h {
		m := (l + h) / 2

		if arr[m] <= e {
			v = arr[m]
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return v
}

// findind smallest element >= x
func ceil(s, e int, arr []int) int {
	v := -1
	l, h := 0, s-1
	for l <= h {
		m := (l + h) / 2

		if arr[m] >= e {
			v = arr[m]
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return v
}
