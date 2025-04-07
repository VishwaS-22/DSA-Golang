package main

import (
	"fmt"
)

func main() {
	//arr := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	// ans := mergeOverlapInt(arr)

	// fmt.Println("The merged intervals are:")

	// for _, v := range ans {
	// 	fmt.Printf("[%d, %d]", v[0], v[1])
	// }
	// fmt.Println()

	// Insert Interval
	arr := [][]int{{1,3},{4,5}}
	new_arr := []int{2,3}
	merged_ans := insertInt(arr, new_arr)
	fmt.Println("The merged intervals after inserting are:")
	for _, v := range merged_ans {
		fmt.Printf("[%d, %d]", v[0], v[1])
	}
	fmt.Println()
}

func sortInt(arr [][]int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j][0] > arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
