package main

func maxArea(arr []int) int {
	left, right := 0, len(arr)-1
	maxA := 0

	for left < right {
		w := right - left
		h := min(arr[left], arr[right])
		area := w * h
		maxA = max(maxA, area)
		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
