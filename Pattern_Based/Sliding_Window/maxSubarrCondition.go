package main

import "fmt"

func maxSubarrCondition(arr []int, k int) int64 {
	fmt.Println("Max sub array sum with k length without Considering repeated vals:")
	var maxSum, windowSum int64 // Use int64 for calculations
	left := 0
	numCount := make(map[int]int) // Track occurrences of elements in the window

	for right := 0; right < len(arr); right++ {
		numCount[arr[right]]++
		windowSum += int64(arr[right]) // Convert to int64

		// If a duplicate is found, shrink the window
		for numCount[arr[right]] > 1 {
			numCount[arr[left]]--
			windowSum -= int64(arr[left]) // Convert to int64
			left++
		}

		// Ensure the window has exactly k distinct elements before considering max sum
		if right-left+1 == k {
			maxSum = max64(maxSum, windowSum)

			// Slide the window forward by removing the left element
			numCount[arr[left]]--
			windowSum -= int64(arr[left]) // Convert to int64
			left++
		}
	}

	return maxSum
}

// Helper function for max
func max64(a, b int64) int64 { // Use int64 to match function return type
	if a > b {
		return a
	}
	return b
}
