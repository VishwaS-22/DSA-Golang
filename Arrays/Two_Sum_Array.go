package main

import (
	"fmt"
	"sort"
)

func two_sum_pair(s, t int, arr []int) {
	// //Brute Force, Yes or No as well as Index with Pairs
	// for i := 0; i < s; i++ {
	// 	for j := i + 1; j < s; j++ {
	// 		if arr[i]+arr[j] == t {
	// 			fmt.Println("Yes, Pair is available")
	// 			fmt.Printf("Index is (%d, %d) and Pair found: (%d, %d)\n", i, j, arr[i], arr[j])
	// 			return // Exit immediately after finding a pair
	// 		}
	// 	}
	// }
	// fmt.Println("No, Pair is not available")

	// //Better Solution

	// hashmap := make(map[int]bool)

	// for i := 0; i < s; i++ {
	// 	c := t - arr[i]

	// 	if hashmap[c] {
	// 		fmt.Println("Yes, The pair is available.")
	// 		return
	// 	}

	// 	hashmap[arr[i]] = true
	// }
	// fmt.Println("No, The pair is available.")

	// Optimal Sol
	sort.Ints(arr)

	// Initialize two pointers.
	left, right := 0, len(arr)-1

	// Use the two-pointer technique to find a pair that sums to target.
	for left < right {
		sum := arr[left] + arr[right]
		if sum == t {
			fmt.Printf("Pair found: %d + %d = %d\n", arr[left], arr[right], t)
			return
		} else if sum < t {
			// If the sum is too small, move the left pointer to the right.
			left++
		} else {
			// If the sum is too large, move the right pointer to the left.
			right--
		}
	}

	fmt.Println("No pair found with the given target sum.")

}
