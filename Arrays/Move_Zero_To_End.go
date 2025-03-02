//Move Zero's of an array to the end of the array

package main

func Zero_To_End(arr []int) []int {
	//Brute Force
	// t := []int{}
	// n := len(arr)
	// for i := range arr {
	// 	if arr[i] != 0 {
	// 		t = append(t, arr[i])
	// 	}
	// }

	// tsize := len(t)

	// for i := 0; i < tsize; i++ {
	// 	arr[i] = t[i]
	// }

	// for i := tsize; i < n; i++ {
	// 	arr[i] = 0
	// }
	// return arr

	n := len(arr)
	j := -1
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			j = i
			break
		}
	}
	if j == -1 {
		return arr
	}

	for i := j + 1; i < n; i++ {
		if arr[i] != 0 {
			t := arr[j]
			arr[j] = arr[i]
			arr[i] = t
			j++
		}
	}
	return arr
}
