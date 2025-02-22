package main

func BS_Search(t int, arr []int) int {
	return BS_Recursive(arr, 0, len(arr)-1, t)
}
func BS_Recursive(arr1 []int, l, h, t int) int {
	if l > h {
		return -1
	}

	m := (l + h) / 2
	if arr1[m] == t {
		return m
	} else if t > arr1[m] {
		return BS_Recursive(arr1, m+1, h, t)
	} else {
		return BS_Recursive(arr1, l, m-1, t)
	}
}
