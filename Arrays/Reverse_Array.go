package main

func Array_rev(i, n int, arr []int) {
	if i >= n/2 {
		return
	}
	// Perform the swap directly on the array
	//arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	
	swap(arr, i, n-i-1)
	Array_rev(i+1, n, arr)
}

func swap(arr []int, a, b int) {
	t := arr[a]
	arr[a] = arr[b]
	arr[b] = t
}
