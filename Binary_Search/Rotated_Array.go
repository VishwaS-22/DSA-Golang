package main

func search_rotated_1(s, t int, arr []int) int {
	l, h := 0, s-1
	for l <= h {
		m := (l + h) / 2

		if arr[m] == t {
			return m
		}
		//With edge case
		if arr[l] == arr[m] && arr[m] == arr[l] {
			l++
			h--
			continue
		}
		// if left part is sorted
		if arr[l] <= arr[m] {
			if arr[l] <= t && t <= arr[m] {
				h = m - 1
			} else {
				l = m + 1
			}
		} else {
			if arr[m] <= t && t <= arr[h] {
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}
	return -1
}
