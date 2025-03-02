package main

//import "fmt"

func mergeSort(arr []int, low, high int) {
	if low < high {
		// Find the middle point
		mid := (low + high) / 2
		// Recursively sort the left half
		mergeSort(arr, low, mid)
		// Recursively sort the right half
		mergeSort(arr, mid+1, high)
		// Merge the two sorted halves
		merge(arr, low, mid, high)
	}
}
func merge(arr []int, low, mid, high int) {
	// Create a temporary slice to hold merged results.
	temp := make([]int, 0, high-low+1)

	// Initialize pointers for the left and right subarrays.
	left, right := low, mid+1

	// Merge elements from both halves while both have remaining elements.
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	// Append remaining elements from the left half.
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	// Append remaining elements from the right half.
	for right <= high {
		temp = append(temp, arr[right])
		right++
	}

	// Copy the merged elements from temp back into arr.
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
}

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	l := mergeSort(arr[:mid]) // 0 to before mid index
// 	r := mergeSort(arr[mid:]) // from mid index to n-1
// 	return mergeArr(l, r)
// }

// func mergeArr(l, r []int) []int {
// 	t := []int{}

// 	i, j := 0, 0

// 	for i < len(l) && j < len(r) {
// 		if l[i] <= r[j] {
// 			t = append(t, l[i])
// 			i++
// 		} else {
// 			t = append(t, r[j])
// 			j++
// 		}
// 	}

// 	// Add Remaining Elements
// 	// for i < len(l) {
// 	// 	t = append(t, l[i])
// 	// 	i++
// 	// }
// 	t = append(t, l[i:]...)
// 	// for j < len(r) {
// 	// 	t = append(t, r[j])
// 	// 	j++
// 	// }
// 	t = append(t, r[j:]...)

// 	return t

// }
