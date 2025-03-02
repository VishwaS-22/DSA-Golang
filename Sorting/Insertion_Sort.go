package main

import "fmt"

func Insertion_Sort(arr []int, s int) {
	for i := 0; i <= s-1; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			// t := arr[j-1]
			// arr[j-1] = arr[j]
			// arr[j] = t
			// Swap arr[j] and arr[j-1]
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	fmt.Print("The sorted arr is: ", arr)
}
