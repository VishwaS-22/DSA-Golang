package main

import (
	// "math"
	"fmt"
)

func maxSubarrsum(arr []int, k int) int {
	fmt.Println("Max sub array sum with k length by Considering repeated vals:")
	maxSum, windowSum := 0, 0
	l := 0

	for r := 0; r < len(arr); r++ {
		windowSum += arr[r]

		if r >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[l]
			l++
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
