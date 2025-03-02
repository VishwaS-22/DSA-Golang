package main

func lower_bound(s, x int, arr []int) int {
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

func upper_bound(s, x int, arr []int) int {
	v := s //6
	l, h := 0, s-1 //0, 5
	for l <= h { // 0<=5, 3<=5. 4<=5, 5>]5
		m := (l + h) / 2 //2 //4 //4 //5

		if arr[m] > x { // 3>6, No , 5>6,  No, 5>6 No,
			v = m
			h = m - 1
		} else {
			l = m + 1 //3rd index //4th index //5th index //6th index
		}
	}
	return v
}
