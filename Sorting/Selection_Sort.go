package main

import "fmt"

func Selection_Sort(s int, arr []int) {
	for i := 0; i < s-2; i++ {
		min := i
		for j := i; j < s-1; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		t := arr[min]
		arr[min] = arr[i]
		arr[i] = t
	}

	fmt.Println("The sorted arr is: ", arr)
}
