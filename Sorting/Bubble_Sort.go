package main

import "fmt"

func Bubble_Sort(s int, arr []int) {
	for i := s - 1; i >= 1; i-- {
		c := 0
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
				c = 1
			}
		}
		if c == 0 {
			break
		}
	}
	fmt.Print("The sorted arr is: ", arr)
}
