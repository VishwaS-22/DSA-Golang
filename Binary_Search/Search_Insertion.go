package main

// Search in an arr and find the index where the i/p element fits in.
func Search_Insertion(s, x int, arr []int) int {
	v := s
	l, h := 0, s-1
	for l <= h {
		m := (l + h) / 2

		if arr[m] >= x {
			v = m
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return v
}
