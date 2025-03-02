package main

import "fmt"

func IsSorted(arr3 []int) {
	isSorted := true
	for i := range arr3 {
		if i == 0 {
			continue
		}
		if arr3[i] < arr3[i-1] {
			isSorted = false
			break
		}
	}
	if isSorted {
		fmt.Println("Sorted One")
	} else {
		fmt.Println("Not a Sorted One")
	}
}
