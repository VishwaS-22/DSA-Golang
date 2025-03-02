package main

func Left_Rotate(arr []int) []int {
	t := arr[0]
	n := len(arr)
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i]
	}
	arr[n-1] = t

	return arr
}
