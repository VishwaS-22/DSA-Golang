package main

import "math"

func SecondSmallest(arr2 []int) int {
	smallest := arr2[0]
	secSmallest := math.MaxInt

	for _, v := range arr2[1:] {
		if v < smallest {
			secSmallest = smallest
			smallest = v
		} else if v != smallest && v < secSmallest {
			secSmallest = v
		}
	}
	return secSmallest
}
