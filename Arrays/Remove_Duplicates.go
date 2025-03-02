// Removing duplicate elements from a Sorted array.

package main

// {1,1,2,2,2,3,3}

func RemoveDup(arr4 []int) int {
	i := 0
	for j := range arr4[1:] {
		if arr4[j] != arr4[i] {
			arr4[i+1] = arr4[j]
			i++
		}
	}
	return i + 1
}
