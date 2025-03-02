package main

///import "fmt"

func Square_root(n int) int {
	l, h := 1, n

	for l <= h {
		m := (l + h) / 2
		v := m * m

		if v <= n {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return h
}
