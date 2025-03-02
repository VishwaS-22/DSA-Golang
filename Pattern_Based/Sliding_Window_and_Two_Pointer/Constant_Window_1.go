//Max sum of consecutive k elements in array

package main

import (
	"fmt"
)

var k int = 4

func main() {
	var arr = []int{-1, 1, 2, 3, 4, 5, 6, 7}
	sum := 0

	for s := 0; s < k; s++ {
		sum += arr[s] //7
		//fmt.Printf("The first k sum is %v\n", sum)
	}
	maxSum := sum
	s := 0
	e := k - 1 // 3, 4
	n := len(arr)

	for e < n-1 {
		sum = sum - arr[s] + arr[e+1] //6+4
		s++
		e++
		if sum > maxSum {
			maxSum = sum
		}
	}

	fmt.Printf("The max sum of array in k consecutive is %v\n", maxSum)

	var nums = []int{6, 2, 3, 4, 7, 2, 1, 7, 1}
	k := 4

	fmt.Printf("The max points of n cards is %v\n", Max_Points(nums, k))
}
